package authen

import (
	"github.com/joyqi/go-feishu/api"
	"net/http"
)

const (
	UserInfoURL = "/authen/v1/user_info"
)

// UserInfoData represents the response data of UserInfo
type UserInfoData struct {
	Name            string `json:"name"`
	EnName          string `json:"en_name"`
	AvatarURL       string `json:"avatar_url"`
	AvatarThumb     string `json:"avatar_thumb"`
	AvatarMiddle    string `json:"avatar_middle"`
	AvatarBig       string `json:"avatar_big"`
	OpenId          string `json:"open_id"`
	UnionId         string `json:"union_id"`
	Email           string `json:"email"`
	EnterpriseEmail string `json:"enterprise_email"`
	UserId          string `json:"user_id"`
	Mobile          string `json:"mobile"`
	TenantKey       string `json:"tenant_key"`
}

type UserInfo api.Api

// Get fetches the user info through the access token.
func (a *UserInfo) Get() (data *UserInfoData, err error) {
	return api.MakeApi[UserInfoData](a.Client, http.MethodGet, UserInfoURL, nil)
}
