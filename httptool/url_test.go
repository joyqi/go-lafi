package httptool

import (
	"net/url"
	"testing"
)

func TestMakeURL(t *testing.T) {
	target := "https://example.com?state=production"
	source := MakeURL("https://example.com", url.Values{"state": {"production"}})

	if target != source {
		t.Fail()
	}
}

func TestMakeStructureURL(t *testing.T) {
	type testParams struct {
		State    string `url:"state"`
		UserName string `url:"user_name"`
	}

	target := "https://example.com/?state=test&user_name=hello"
	source := MakeStructureURL("https://example.com/", &testParams{
		State:    "test",
		UserName: "hello",
	})

	if target != source {
		t.Fail()
	}
}

func TestMakeTemplateURL(t *testing.T) {
	target := "https://example.com/123"
	source := MakeTemplateURL("https://example.com/:id", map[string]string{"id": "123"})

	if target != source {
		t.Fail()
	}
}
