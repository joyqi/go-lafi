package auth

import "github.com/joyqi/go-feishu/api"

const (
	AppCommonURL   = "/auth/v3/app_access_token"
	AppInternalURL = "/auth/v3/app_access_token/internal"
)

// AppCommonBody represents a request to retrieve an app token
type AppCommonBody struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	AppTicket string `json:"app_ticket"`
}

// AppInternalBody represents a request to retrieve an app token
type AppInternalBody struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type App api.Api

// CommonAccessToken retrieves a common token from the app token endpoint
func (a *App) CommonAccessToken(body *AppCommonBody) (string, int64, error) {
	return MakeTokenApi(a, "app_access_token", AppCommonURL, body)
}

// InternalAccessToken retrieves an internal token from the app token endpoint
func (a *App) InternalAccessToken(body *AppInternalBody) (string, int64, error) {
	return MakeTokenApi(a, "app_access_token", AppInternalURL, body)
}
