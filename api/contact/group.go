package contact

import "github.com/joyqi/go-oauth2-feishu/api"

type Group interface {
	MemberBelong(openId string)
}

type GroupApi struct {
	api.Api
}

func (a *GroupApi) MemberBelong(openId string) {
	a.Client.Get()
}
