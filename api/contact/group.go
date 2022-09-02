package contact

import (
	"github.com/joyqi/go-oauth2-feishu/api"
	"net/http"
)

var (
	GroupMemberBelongURL = "https://open.feishu.cn/open-apis/contact/v3/group/member_belong"
)

type MemberBelongData struct {
	GroupList []string `json:"group_list,flow"`
	HasMore   bool     `json:"has_more"`
}

type GroupApi interface {
	MemberBelong(openId string) (*api.Response[MemberBelongData], error)
}

type Group struct {
	api.Api
}

func (a *Group) MemberBelong(openId string) (*MemberBelongData, error) {
	url := api.MakeURL(GroupMemberBelongURL, api.Param{
		Key:   "member_id",
		Value: openId,
	})
	return api.MakeApi[MemberBelongData](a.Api.Client, http.MethodGet, url, nil)
}
