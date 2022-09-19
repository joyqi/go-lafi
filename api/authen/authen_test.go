package authen

import (
	"context"
	"github.com/joyqi/go-feishu/oauth2"
	"os"
	"testing"
	"time"
)

var conf = &oauth2.Config{
	AppID:       os.Getenv("APP_ID"),
	AppSecret:   os.Getenv("APP_SECRET"),
	RedirectURL: "https://example.com",
}

var token = &oauth2.Token{
	AccessToken:  os.Getenv("ACCESS_TOKEN"),
	RefreshToken: os.Getenv("REFRESH_TOKEN"),
	Expiry:       time.Now().Truncate(time.Hour),
}

func TestAuthen_UserInfo(t *testing.T) {
	ts := conf.TokenSource(context.Background(), token)
	api := &Authen{Client: ts.Client()}
	_, err := api.UserInfo()

	if err != nil {
		t.Error(err)
	}
}
