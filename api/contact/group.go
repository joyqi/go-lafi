package contact

import (
	"github.com/joyqi/go-feishu/api"
	"github.com/joyqi/go-feishu/httptool"
	"net/http"
	"net/url"
)

var (
	GroupMemberBelongURL = "https://open.feishu.cn/open-apis/contact/v3/group/member_belong"
)

type MemberBelongData struct {
	GroupList []string `json:"group_list,flow"`
	HasMore   bool     `json:"has_more"`
}

type GroupApi interface {
	MemberBelong(openId string) (*MemberBelongData, error)
}

type Group struct {
	api.Api
}

func (a *Group) MemberBelong(openId string) (*MemberBelongData, error) {
	u := httptool.MakeURL(GroupMemberBelongURL, url.Values{"member_id": {openId}})
	return api.MakeApi[MemberBelongData](a.Api.Client, http.MethodGet, u, nil)
}
