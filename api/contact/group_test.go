package contact

import (
	"context"
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

	_, err := a.MemberBelong(&GroupMemberBelongParams{
		MemberId:     "9a4386bc",
		MemberIdType: "user_id",
	})

	if err != nil {
		t.Error(err)
	}
}
