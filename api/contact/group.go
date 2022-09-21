package contact

import (
	"github.com/creasty/defaults"
	"github.com/joyqi/go-lafi/api"
	"github.com/joyqi/go-lafi/httptool"
	"net/http"
)

const (
	GroupURL             = "/contact/v3/group/:group_id"
	GroupCreateURL       = "/contact/v3/group"
	GroupSimpleListURL   = "/contact/v3/group/simplelist"
	GroupMemberBelongURL = "/contact/v3/group/member_belong"
)

// GroupType represents the type of group.
// GroupTypeNormal is the normal group.
// GroupTypeDynamic is the dynamic group.
type GroupType int8

const (
	GroupTypeNormal GroupType = 1 + iota
	GroupTypeDynamic
)

// GroupCreateBody represents the params of Group.Create
type GroupCreateBody struct {
	GroupId     string    `json:"group_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        GroupType `json:"type" default:"1"`
}

// GroupCreateData represents the response data of Group.Create
type GroupCreateData struct {
	GroupId string `json:"group_id"`
}

// GroupPatchBody represents the params of Group.Patch
type GroupPatchBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GroupPatchData represents the response data of Group.Patch
type GroupPatchData = api.EmptyData

// GroupDeleteData represents the response data of Group.Delete
type GroupDeleteData = api.EmptyData

// GroupGetData represents the response data of Group.Get
type GroupGetData struct {
	Group struct {
		GroupId               string `json:"group_id"`
		Name                  string `json:"name"`
		Description           string `json:"description"`
		MemberUserCount       int    `json:"member_user_count"`
		MemberDepartmentCount int    `json:"member_department_count"`
	} `json:"group"`
}

// GroupSimpleListParams represents the params of Group.SimpleList
type GroupSimpleListParams struct {
	PageSize  int       `url:"page_size" default:"50"`
	PageToken string    `url:"page_token"`
	Type      GroupType `url:"type" default:"1"`
}

// GroupSimpleListData represents the response data of Group.SimpleList
type GroupSimpleListData struct {
	GroupList []struct {
		Id                    string `json:"id"`
		Name                  string `json:"name"`
		Description           string `json:"description"`
		MemberUserCount       int    `json:"member_user_count"`
		MemberDepartmentCount int    `json:"member_department_count"`
	} `json:"grouplist"`

	PageToken string `json:"page_token"`
	HasMore   bool   `json:"has_more"`
}

// GroupMemberBelongParams represents the params of Group.MemberBelong
type GroupMemberBelongParams struct {
	MemberId     string    `url:"member_id"`
	MemberIdType string    `url:"member_id_type" default:"open_id"`
	GroupType    GroupType `url:"group_type" default:"1"`
	PageSize     int       `url:"page_size" default:"500"`
	PageToken    string    `url:"page_token"`
}

// GroupMemberBelongData represents the response data of Group.MemberBelong
type GroupMemberBelongData struct {
	GroupList []string `json:"group_list"`
	HasMore   bool     `json:"has_more"`
	PageToken string   `json:"page_token"`
}

type Group api.Api

// Create creates a group by given GroupCreateBody.
// See https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/create for more details.
func (g *Group) Create(body *GroupCreateBody) (*GroupCreateData, error) {
	if err := defaults.Set(body); err != nil {
		return nil, err
	}

	return api.MakeApi[GroupCreateData](g.Client, http.MethodPost, GroupCreateURL, body)
}

// Delete deletes a group through group_id.
// See https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/delete for more details.
func (g *Group) Delete(groupId string) (*GroupDeleteData, error) {
	url := httptool.MakeTemplateURL(GroupURL, map[string]string{"group_id": groupId})
	return api.MakeApi[GroupDeleteData](g.Client, http.MethodDelete, url, nil)
}

// Get retrieves the group information through group_id.
// See https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/get for more details.
func (g *Group) Get(groupId string) (*GroupGetData, error) {
	url := httptool.MakeTemplateURL(GroupURL, map[string]string{"group_id": groupId})
	return api.MakeApi[GroupGetData](g.Client, http.MethodGet, url, nil)
}

// Patch updates the specified group information by given GroupPatchBody.
// See https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/patch for more details.
func (g *Group) Patch(groupId string, body *GroupPatchBody) (*GroupPatchData, error) {
	url := httptool.MakeTemplateURL(GroupURL, map[string]string{"group_id": groupId})
	return api.MakeApi[GroupPatchData](g.Client, http.MethodPatch, url, body)
}

// SimpleList retrieves the group list through given GroupSimpleListParams.
// See https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/simplelist for more details.
func (g *Group) SimpleList(params *GroupSimpleListParams) (*GroupSimpleListData, error) {
	if err := defaults.Set(params); err != nil {
		return nil, err
	}

	url := httptool.MakeStructureURL(GroupSimpleListURL, params)
	return api.MakeApi[GroupSimpleListData](g.Client, http.MethodGet, url, params)
}

// MemberBelong returns the groups that the member belongs to.
// See https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/member_belong for more details.
func (g *Group) MemberBelong(params *GroupMemberBelongParams) (*GroupMemberBelongData, error) {
	if err := defaults.Set(params); err != nil {
		return nil, err
	}

	url := httptool.MakeStructureURL(GroupMemberBelongURL, params)
	return api.MakeApi[GroupMemberBelongData](g.Client, http.MethodGet, url, nil)
}
