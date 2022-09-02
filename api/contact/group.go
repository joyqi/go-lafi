package contact

import (
	"github.com/creasty/defaults"
	"github.com/joyqi/go-feishu/api"
	"github.com/joyqi/go-feishu/httptool"
	"net/http"
)

var (
	GroupMemberBelongURL = "https://open.feishu.cn/open-apis/contact/v3/group/member_belong"
)

type GroupMemberBelongParams struct {
	MemberId     string `url:"member_id"`
	MemberIdType string `url:"member_id_type" default:"open_id"`
	GroupType    int    `url:"group_type" default:"1"`
	PageSize     int    `url:"page_size" default:"500"`
	PageToken    string `url:"page_token"`
}

type GroupMemberBelongData struct {
	GroupList []string `json:"group_list,flow"`
	HasMore   bool     `json:"has_more"`
}

type GroupApi interface {
	MemberBelong(openId string) (*GroupMemberBelongData, error)
}

type Group struct {
	api.Api
}

func (a *Group) MemberBelong(params *GroupMemberBelongParams) (*GroupMemberBelongData, error) {
	if err := defaults.Set(params); err != nil {
		return nil, err
	}

	u := httptool.MakeStructureURL(GroupMemberBelongURL, params)
	return api.MakeApi[GroupMemberBelongData](a.Api.Client, http.MethodGet, u, nil)
}
