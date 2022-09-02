package contact

import (
	"context"
	"github.com/joyqi/go-feishu/oauth2"
	"os"
	"testing"
)

var conf = &oauth2.Config{
	AppID:       os.Getenv("APP_ID"),
	AppSecret:   os.Getenv("APP_SECRET"),
	RedirectURL: "https://example.com",
}

func TestGroup_Create(t *testing.T) {
	c := conf.Client(context.Background())
	a := &Group{Client: c}

	data, err := a.Create(&GroupCreateParams{
		GroupId: "test001",
		Name:    "test group",
	})

	if err != nil {
		t.Error(err)
	}

	if data.GroupId != "test001" {
		t.Fail()
	}
}

func TestGroup_Delete(t *testing.T) {
	c := conf.Client(context.Background())
	a := &Group{Client: c}

	_, err := a.Delete("test001")

	if err != nil {
		t.Error(err)
	}
}

func TestGroup_MemberBelong(t *testing.T) {
	c := conf.Client(context.Background())
	a := &Group{Client: c}

	_, err := a.MemberBelong(&GroupMemberBelongParams{
		MemberId:     "9a4386bc",
		MemberIdType: "user_id",
	})

	if err != nil {
		t.Error(err)
	}
}
