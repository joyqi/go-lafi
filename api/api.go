package api

import (
	"errors"
	"github.com/joyqi/go-oauth2-feishu/oauth2"
	"net/url"
)

type Api struct {
	Client *oauth2.Client
}

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type Param struct {
	Key   string
	Value string
}

// MakeApi make a standard Api request and handle the response accordingly
func MakeApi[T any](c *oauth2.Client, method string, uri string, body interface{}) (*T, error) {
	resp := &Response[T]{}
	err := c.Request(method, uri, body, resp)

	if err != nil {
		return nil, err
	} else if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}

	return &resp.Data, nil
}

func MakeURL(uri string, params ...Param) string {
	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}

	query := u.Query()
	for _, param := range params {
		if param.Value != "" {
			query.Set(param.Key, param.Value)
		}
	}

	u.RawQuery = query.Encode()
	return u.String()
}
