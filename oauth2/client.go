package oauth2

import (
	"context"
	"github.com/joyqi/go-oauth2-feishu/http"
)

type MethodRequest = func(uri string, body interface{}, data interface{}) error

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

	header := http.Header{
		Key:   "Authorization",
		Value: "Bearer " + token,
	}

	return http.Request(c.ctx, &http.RequestOptions{
		URI:         uri,
		Method:      method,
		Headers:     []http.Header{header},
		ContentType: "application/json; charset=utf-8",
		JSONBody:    body,
	}, data)
}
