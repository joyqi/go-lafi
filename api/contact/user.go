package contact

import "github.com/joyqi/go-feishu/api"

var (
	UserCreateURL = "https://open.feishu.cn/open-apis/contact/v3/users"
)

type UserIdType string

const (
	UserIdTypeUserId  UserIdType = "user_id"
	UserIdTypeOpenId             = "open_id"
	UserIdTypeUnionId            = "union_id"
)

type UserCreateParams struct {
	UserIdType       UserIdType       `url:"user_id_type" default:"open_id"`
	DepartmentIdType DepartmentIdType `url:"department_id_type" default:"department_id"`
	ClientToken      string           `url:"client_token"`
}

type UserCreateBody struct {
	Name     string             `json:"name"`
	I18nName DepartmentI18nName `json:"i18n_name"`
}

type UserCreateData struct {
}

type User api.Api

func (u *User) Create(params UserCreateParams, body UserCreateBody) (UserCreateData, error) {
}
