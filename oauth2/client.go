package oauth2

import (
	"context"
	"github.com/joyqi/go-feishu/httptool"
	"net/http"
)

// Client returns an authorized client for the given endpoint
func (c *Config) Client(ctx context.Context) *Client {
	return &Client{
		ctx:  ctx,
		conf: c,
	}
}

type Client struct {
	ctx  context.Context
	conf *Config
}

func (c *Client) Request(method string, uri string, body interface{}, data interface{}) error {
	token, err := c.conf.TenantToken(c.ctx)
	if err != nil {
		return err
	}

	header := httptool.Header{
		Key:   "Authorization",
		Value: "Bearer " + token,
	}

	return httptool.Request(c.ctx, &httptool.RequestOptions{
		URI:         uri,
		Method:      method,
		Headers:     []httptool.Header{header},
		ContentType: "application/json; charset=utf-8",
		JSONBody:    body,
	}, data)
}

func httpPost(ctx context.Context, uri string, body interface{}, data interface{}, headers ...httptool.Header) error {
	return httptool.Request(ctx, &httptool.RequestOptions{
		URI:         uri,
		Method:      http.MethodPost,
		Headers:     headers,
		ContentType: "application/json; charset=utf-8",
		JSONBody:    body,
	}, data)
}
