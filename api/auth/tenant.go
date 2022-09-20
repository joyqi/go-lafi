package auth

import "github.com/joyqi/go-feishu/api"

const (
	TenantCommonURL   = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token"
	TenantInternalURL = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
)

// TenantCommonBody represents a request to retrieve a tenant token
type TenantCommonBody struct {
	AppAccessToken string `json:"app_access_token"`
	TenantKey      string `json:"tenant_key"`
}

// TenantInternalBody represents a request to retrieve a tenant token
type TenantInternalBody struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type Tenant api.Api

// CommonAccessToken retrieves a common token from the tenant token endpoint
func (t *Tenant) CommonAccessToken(body *TenantCommonBody) (string, int64, error) {
	return MakeTokenApi(t, "tenant_access_token", TenantCommonURL, body)
}

// InternalAccessToken retrieves a internal token from the tenant token endpoint
func (t *Tenant) InternalAccessToken(body *TenantInternalBody) (string, int64, error) {
	return MakeTokenApi(t, "tenant_access_token", TenantInternalURL, body)
}
