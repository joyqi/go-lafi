package contact

import (
	"context"
	"fmt"
	"github.com/joyqi/go-feishu/api"
	"github.com/joyqi/go-feishu/oauth2"
	"os"
	"testing"
)

var conf = &oauth2.Config{
	AppID:       os.Getenv("APP_ID"),
	AppSecret:   os.Getenv("APP_SECRET"),
	RedirectURL: "https://example.com",
}

func TestGroup_MemberBelong(t *testing.T) {
	c := conf.Client(context.Background())
	a := &Group{Api: api.Api{Client: c}}

	data, err := a.MemberBelong("9a4386bc")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data)
}
