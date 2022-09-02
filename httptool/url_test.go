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
