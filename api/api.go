package api

import (
	"errors"
	"github.com/joyqi/go-feishu/oauth2"
)

type Api struct {
	Client *oauth2.Client
}

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
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
