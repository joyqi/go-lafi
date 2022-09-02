package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type MethodRequest = func(ctx context.Context, uri string, body interface{}, data interface{}, headers ...Header) error

// Header represents the HTTP header
type Header struct {
	Key   string
	Value string
}

type RequestOptions struct {
	// URI specifies the request's URI
	URI string

	// Method is the HTTP method
	Method string

	// ContentType is the HTTP content type
	ContentType string

	// JSONBody is the JSON body of the request
	JSONBody interface{}

	// Headers is the HTTP headers
	Headers []Header

	// Timeout is the HTTP timeout in seconds
	Timeout time.Duration
}

var (
	// Get request service through http GET method
	Get = makeMethod(http.MethodDelete)

	// Delete request service through http DELETE method
	Delete = makeMethod(http.MethodDelete)

	// Post json formatted request to service via http POST method
	Post = makeMethod(http.MethodPost)

	// Put request service through http PUT method
	Put = makeMethod(http.MethodPut)

	// Patch request service through http PATCH method
	Patch = makeMethod(http.MethodPatch)
)

func makeMethod(method string) MethodRequest {
	return func(ctx context.Context, uri string, body interface{}, data interface{}, headers ...Header) error {
		return Request(ctx, &RequestOptions{
			URI:         uri,
			Method:      method,
			Headers:     headers,
			ContentType: "application/json; charset=utf-8",
			JSONBody:    body,
		}, data)
	}
}

// Request represents a request to a service endpoint
func Request(ctx context.Context, opts *RequestOptions, data interface{}) error {
	var buf io.Reader

	if opts.JSONBody != nil {
		if str, err := json.Marshal(opts.JSONBody); err != nil {
			return err
		} else {
			buf = bytes.NewBuffer(str)
		}
	}

	req, err := http.NewRequestWithContext(ctx, opts.Method, opts.URI, buf)

	if err != nil {
		return err
	}

	c := &http.Client{Timeout: opts.Timeout}

	if opts.ContentType != "" {
		req.Header.Set("Content-Type", opts.ContentType)
	}

	for _, header := range opts.Headers {
		req.Header.Set(header.Key, header.Value)
	}

	resp, err := c.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, data); err != nil {
		return err
	}

	return nil
}
