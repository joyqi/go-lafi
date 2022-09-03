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

	data, err := a.Create(&GroupCreateBody{
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

func TestGroup_Get(t *testing.T) {
	c := conf.Client(context.Background())
	a := &Group{Client: c}

	g, err := a.Get("test001")

	if err != nil {
		t.Error(err)
	} else if g.Group.Name != "test group" {
		t.Fail()
	}
}

func TestGroup_Patch(t *testing.T) {
	c := conf.Client(context.Background())
	a := &Group{Client: c}

	_, err := a.Patch("test001", &GroupPatchBody{
		Name: "test group 2",
	})

	if err != nil {
		t.Error(err)
	}

	g, err := a.Get("test001")

	if err != nil {
		t.Error(err)
	} else if g.Group.Name != "test group 2" {
		t.Fail()
	}
}

func TestGroup_SimpleList(t *testing.T) {
	c := conf.Client(context.Background())
	a := &Group{Client: c}

	g, err := a.SimpleList(&GroupSimpleListParams{
		PageSize: 10,
	})

	if err != nil {
		t.Error(err)
	} else if len(g.GroupList) == 0 {
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
		MemberIdType: UserIdTypeUserId,
	})

	if err != nil {
		t.Error(err)
	}
}
