package oauth2

import (
	"context"
	"github.com/joyqi/go-feishu/httptool"
	"net/http"
)

// Client returns http client with an authorized token.
func (c *Config) Client(ctx context.Context) *Client {
	return &Client{
		ctx: ctx,
		ts:  c,
	}
}

// ClientTokenSource represents a token source that returns a client token.
type ClientTokenSource interface {
	ClientToken(ctx context.Context) (string, error)
}

// A Client represents a http client with an authorized token.
type Client struct {
	ctx context.Context
	ts  ClientTokenSource
}

// Request performs a http request to the given endpoint with the authorized token.
func (c *Client) Request(method string, uri string, body interface{}, data interface{}) error {
	token, err := c.ts.ClientToken(c.ctx)
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
