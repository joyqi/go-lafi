package contact

import (
	"github.com/creasty/defaults"
	"github.com/joyqi/go-feishu/api"
	"github.com/joyqi/go-feishu/httptool"
	"net/http"
)

var (
	GroupCreateURL       = "https://open.feishu.cn/open-apis/contact/v3/group"
	GroupDeleteURL       = "https://open.feishu.cn/open-apis/contact/v3/group/:group_id"
	GroupMemberBelongURL = "https://open.feishu.cn/open-apis/contact/v3/group/member_belong"
)

type GroupCreateParams struct {
	GroupId     string `json:"group_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        int    `json:"type" default:"1"`
}

type GroupCreateData struct {
	GroupId string `json:"group_id"`
}

type GroupDeleteData struct {
}

type GroupMemberBelongParams struct {
	MemberId     string     `url:"member_id"`
	MemberIdType UserIdType `url:"member_id_type" default:"open_id"`
	GroupType    int        `url:"group_type" default:"1"`
	PageSize     int        `url:"page_size" default:"500"`
	PageToken    string     `url:"page_token"`
}

type GroupMemberBelongData struct {
	GroupList []string `json:"group_list,flow"`
	HasMore   bool     `json:"has_more"`
}

type Group api.Api

func (g *Group) Create(params *GroupCreateParams) (*GroupCreateData, error) {
	if err := defaults.Set(params); err != nil {
		return nil, err
	}

	return api.MakeApi[GroupCreateData](g.Client, http.MethodPost, GroupCreateURL, params)
}

func (g *Group) Delete(groupId string) (*GroupDeleteData, error) {
	u := httptool.MakeTemplateURL(GroupDeleteURL, map[string]string{"group_id": groupId})
	return api.MakeApi[GroupDeleteData](g.Client, http.MethodDelete, u, nil)
}

func (g *Group) MemberBelong(params *GroupMemberBelongParams) (*GroupMemberBelongData, error) {
	if err := defaults.Set(params); err != nil {
		return nil, err
	}

	u := httptool.MakeStructureURL(GroupMemberBelongURL, params)
	return api.MakeApi[GroupMemberBelongData](g.Client, http.MethodGet, u, nil)
}
