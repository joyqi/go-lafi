package httptool

import (
	"context"
	"net"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

type testResponse struct {
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type testJsonResponse struct {
	Json testRequest `json:"json"`
}

type testRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Count int    `json:"count"`
}

var req = testRequest{
	Name:  "test name",
	Value: "test value",
	Count: 1,
}

func TestGet(t *testing.T) {
	resp := &testResponse{}
	err := Request(context.Background(), &RequestOptions{
		URI:    "https://httpbin.org/get",
		Method: http.MethodGet,
	}, resp)

	if err != nil {
		t.Error(err)
	}
}

func TestJsonResponse(t *testing.T) {
	methods := []string{http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch}

	for _, method := range methods {
		resp := &testJsonResponse{}
		err := Request(context.Background(), &RequestOptions{
			URI:      "https://httpbin.org/" + strings.ToLower(method),
			Method:   method,
			JSONBody: &req,
		}, resp)

		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(req, resp.Json) {
			t.Fail()
		}
	}
}

func TestTimeout(t *testing.T) {
	resp := &testResponse{}
	err := Request(context.Background(), &RequestOptions{
		URI:     "https://httpbin.org/delay/3",
		Method:  http.MethodGet,
		Timeout: 2 * time.Second,
	}, resp)

	if err == nil {
		t.Fail()
	}

	if err, ok := err.(net.Error); ok && err.Timeout() {
		t.SkipNow()
	}

	t.Fail()
}
