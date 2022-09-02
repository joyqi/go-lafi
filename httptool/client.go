package httptool

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
