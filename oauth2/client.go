package oauth2

import (
	"context"
	"github.com/joyqi/go-feishu/api"
	"github.com/joyqi/go-feishu/httptool"
)

// ClientSource represents the source which provides a Client method to retrieve a Client.
type ClientSource interface {
	TokenSource
	Client() api.Client
}

// simpleClient is a client that does not use a token.
type simpleClient struct {
	ctx context.Context
	t   Type
}

// Request performs a http request to the given endpoint without the authorized token.
func (c *simpleClient) Request(method string, uri string, body interface{}, data interface{}) error {
	return httptool.Request(c.ctx, &httptool.RequestOptions{
		URI:         adjustURL(c.t, uri),
		Method:      method,
		ContentType: "application/json; charset=utf-8",
		JSONBody:    body,
	}, data)
}

// tokenClient represents a client that performs requests with an authorized token.
type tokenClient struct {
	ctx context.Context
	ts  TokenSource
	t   Type
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
		URI:         adjustURL(c.t, uri),
		Method:      method,
		Headers:     []httptool.Header{header},
		ContentType: "application/json; charset=utf-8",
		JSONBody:    body,
	}, data)
}

// adjustURL adjusts the url based on the type of the client.
// For TypeFeishu, the url will be prefixed with the Feishu API URL.
// For TypeLark, the url will be prefixed with the Lark API URL.
func adjustURL(t Type, url string) string {
	prefix := ""

	switch t {
	case TypeFeishu:
		prefix = "https://open.feishu.cn/open-apis"
	case TypeLark:
		prefix = "https://open.larksuite.com/open-apis"
	}

	return prefix + url
}
