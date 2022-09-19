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

var client = conf.TenantTokenSource(context.Background()).Client()

func TestGroup_Create(t *testing.T) {
	a := &Group{Client: client}

	data, err := a.Create(&GroupCreateBody{
		GroupId: "test001",
		Name:    "test group",
	})

	if err != nil {
		t.Error(err)
	} else if data.GroupId != "test001" {
		t.Fail()
	}
}

func TestGroup_Get(t *testing.T) {
	a := &Group{Client: client}

	g, err := a.Get("test001")

	if err != nil {
		t.Error(err)
	} else if g.Group.Name != "test group" {
		t.Fail()
	}
}

func TestGroup_Patch(t *testing.T) {
	a := &Group{Client: client}

	_, err := a.Patch("test001", &GroupPatchBody{
		Name: "test group 2",
	})

	if err != nil {
		t.Fatal(err)
	}

	g, err := a.Get("test001")

	if err != nil {
		t.Error(err)
	} else if g.Group.Name != "test group 2" {
		t.Fail()
	}
}

func TestGroup_SimpleList(t *testing.T) {
	a := &Group{Client: client}

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
	a := &Group{Client: client}

	_, err := a.Delete("test001")

	if err != nil {
		t.Error(err)
	}
}

func TestGroup_MemberBelong(t *testing.T) {
	a := &Group{Client: client}

	_, err := a.MemberBelong(&GroupMemberBelongParams{
		MemberId:     "9a4386bc",
		MemberIdType: UserIdTypeUserId,
	})

	if err != nil {
		t.Error(err)
	}
}
