package oauth2

import (
	"context"
	"os"
	"testing"
)

var conf = &Config{
	AppID:       os.Getenv("FEISHU_APP_ID"),
	AppSecret:   os.Getenv("FEISHU_APP_SECRET"),
	RedirectURL: "https://example.com",
}

func TestConfig_AuthCodeURL(t *testing.T) {
	target := EndpointURL.AuthURL + "?app_id=" + conf.AppID +
		"&redirect_uri=https%3A%2F%2Fexample.com&response_type=code&state=test"
	if conf.AuthCodeURL("test") != target {
		t.Fail()
	}
}

func TestConfig_TenantToken(t *testing.T) {
	_, err := conf.TenantToken(context.Background())

	if err != nil {
		t.Error(err)
	}
}
