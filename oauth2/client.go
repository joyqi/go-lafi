package oauth2

import (
	"context"
	"github.com/joyqi/go-feishu/httptool"
	"net/http"
)

// A Client represents a http client with an authorized token.
type tokenClient struct {
	ctx context.Context
	ts  TokenSource
}

// Request performs a http request to the given endpoint with the authorized token.
func (c *tokenClient) Request(method string, uri string, body interface{}, data interface{}) error {
	token, err := c.ts.Token()
	if err != nil {
		return err
	}

	header := httptool.Header{
		Key:   "Authorization",
		Value: "Bearer " + token.AccessToken,
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
