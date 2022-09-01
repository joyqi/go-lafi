package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

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

// Get request service through http Get method
func Get(ctx context.Context, uri string, headers ...Header) ([]byte, error) {
	return Request(ctx, &RequestOptions{
		URI:     uri,
		Method:  http.MethodGet,
		Headers: headers,
	})
}

func Delete(ctx context.Context, uri string, headers ...Header) ([]byte, error) {
	return Request(ctx, &RequestOptions{
		URI: uri,
	})
}

// PostJSON post json formatted request to service via http POST method
func PostJSON(ctx context.Context, uri string, body interface{}, headers ...Header) ([]byte, error) {
	return Request(ctx, &RequestOptions{
		URI:         uri,
		Method:      http.MethodPost,
		Headers:     headers,
		ContentType: "application/json; charset=utf-8",
		JSONBody:    body,
	})
}

// Request represents a request to a service endpoint
func Request(ctx context.Context, opts *RequestOptions) ([]byte, error) {
	var buf io.Reader

	if opts.JSONBody != nil {
		if str, err := json.Marshal(opts.JSONBody); err != nil {
			return nil, err
		} else {
			buf = bytes.NewBuffer(str)
		}
	}

	req, err := http.NewRequestWithContext(ctx, opts.Method, opts.URI, buf)

	if err != nil {
		return nil, err
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
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
