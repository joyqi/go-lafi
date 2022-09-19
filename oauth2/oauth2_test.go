package oauth2

import (
	"context"
	"os"
	"testing"
	"time"
)

var conf = &Config{
	AppID:       os.Getenv("APP_ID"),
	AppSecret:   os.Getenv("APP_SECRET"),
	RedirectURL: "https://example.com",
}

func TestConfig_AuthCodeURL(t *testing.T) {
	target := AuthURL + "?app_id=" + conf.AppID +
		"&redirect_uri=https%3A%2F%2Fexample.com&response_type=code&state=test"
	if conf.AuthCodeURL("test") != target {
		t.Fail()
	}
}

func TestConfig_TenantToken(t *testing.T) {
	ts := conf.TenantTokenSource(context.Background())
	_, err := ts.Token()

	if err != nil {
		t.Error(err)
	}
}

func TestReuseTokenSource_Token(t *testing.T) {
	tk := &Token{
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
		Expiry:       time.Now().Truncate(time.Hour),
	}

	ts := conf.TokenSource(context.Background(), tk)
	nt, err := ts.Token()

	if err != nil {
		t.Error(err)
	} else if tk.AccessToken == nt.AccessToken {
		t.Fail()
	}
}
