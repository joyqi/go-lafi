package contact

import (
	"github.com/creasty/defaults"
	"github.com/joyqi/go-feishu/api"
	"github.com/joyqi/go-feishu/httptool"
	"net/http"
)

const (
	UserCreateURL = "https://open.feishu.cn/open-apis/contact/v3/users"
)

const (
	UserIdTypeUserId  = "user_id"
	UserIdTypeOpenId  = "open_id"
	UserIdTypeUnionId = "union_id"
)

type UserCreateParams struct {
	UserIdType       string `url:"user_id_type" default:"open_id"`
	DepartmentIdType string `url:"department_id_type" default:"department_id"`
	ClientToken      string `url:"client_token"`
}

type UserCreateBody struct {
	Name     string             `json:"name"`
	I18nName DepartmentI18nName `json:"i18n_name"`
}

type UserCreateData struct {
}

type User api.Api

// Create creates a user by given UserCreateParams and UserCreateBody.
// See https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create for more details.
func (u *User) Create(params *UserCreateParams, body *UserCreateBody) (*UserCreateData, error) {
	if err := defaults.Set(&params); err != nil {
		return nil, err
	}

	url := httptool.MakeStructureURL(UserCreateURL, params)
	return api.MakeApi[UserCreateData](u.Client, http.MethodPost, url, body)
}
