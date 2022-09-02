package oauth2

import (
	"context"
	"fmt"
	"testing"
)

var conf = &Config{
	AppID:       "cli_slkdjalasdkjasd",
	AppSecret:   "dskLLdkasdjlasdKK",
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
	token, err := conf.TenantToken(context.Background())

	if err != nil {
		t.Error(err)
	}

	fmt.Println(token)
}
