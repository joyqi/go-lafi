package oauth2

import (
	"context"
	"github.com/joyqi/go-oauth2-feishu/http"
)

type MethodRequest = func(uri string, body interface{}, data interface{}) error

func (c *Config) Client(ctx context.Context) *Client {
	return &Client{
		Get:    makeMethod(ctx, http.MethodGet, c),
		Post:   makeMethod(ctx, http.MethodPost, c),
		Delete: makeMethod(ctx, http.MethodDelete, c),
		Put:    makeMethod(ctx, http.MethodPut, c),
		Patch:  makeMethod(ctx, http.MethodPatch, c),
	}
}

type Client struct {
	Get    MethodRequest
	Post   MethodRequest
	Delete MethodRequest
	Put    MethodRequest
	Patch  MethodRequest
}

func makeMethod(ctx context.Context, method string, c *Config) MethodRequest {
	return func(uri string, body interface{}, data interface{}) error {
		token, err := c.TenantToken(ctx)
		if err != nil {
			return err
		}

		header := http.Header{
			Key:   "Authorization",
			Value: "Bearer " + token,
		}

		return http.Request(ctx, &http.RequestOptions{
			URI:         uri,
			Method:      method,
			Headers:     []http.Header{header},
			ContentType: "application/json; charset=utf-8",
			JSONBody:    body,
		}, data)
	}
}
