package api

import (
	"errors"
)

type Api struct {
	Client
}

type EmptyData struct {
}

// Client defines the interface of api client
type Client interface {
	Request(method string, uri string, body interface{}, data interface{}) error
}

// Response represents the response data of a standard Api request
type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// MakeApi make a standard Api request and handle the response accordingly
func MakeApi[T any](c Client, method string, uri string, body interface{}) (*T, error) {
	resp := &Response[T]{}
	err := c.Request(method, uri, body, resp)

	if err != nil {
		return nil, err
	} else if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}

	return &resp.Data, nil
}
