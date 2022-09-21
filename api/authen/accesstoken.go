package authen

import (
	"github.com/joyqi/go-lafi/api"
	"net/http"
)

const (
	AccessTokenURL        = "/authen/v1/access_token"
	AccessTokenRefreshURL = "/authen/v1/refresh_access_token"
)

// AccessTokenCreateBody represents the request body of creating AccessToken
type AccessTokenCreateBody struct {
	GrantType string `json:"grant_type"`
	Code      string `json:"code"`
}

// AccessTokenData represents the data of creating AccessToken
type AccessTokenData struct {
	// OpenId represents the open ID of the user
	OpenId string `json:"open_id"`

	// AccessToken is the token used to access the application
	AccessToken string `json:"access_token"`

	// RefreshToken is the token used to refresh the user's access token
	RefreshToken string `json:"refresh_token"`

	// ExpiresIn is the number of seconds the token will be valid
	ExpiresIn int64 `json:"expires_in"`
}

// AccessTokenRefreshBody represents the request body of refreshing AccessToken
type AccessTokenRefreshBody struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type AccessToken api.Api

// Create creates the access token.
func (a *AccessToken) Create(body *AccessTokenCreateBody) (*AccessTokenData, error) {
	return api.MakeApi[AccessTokenData](a.Client, http.MethodPost, AccessTokenURL, body)
}

// Refresh refreshes the access token.
func (a *AccessToken) Refresh(body *AccessTokenRefreshBody) (*AccessTokenData, error) {
	return api.MakeApi[AccessTokenData](a.Client, http.MethodPost, AccessTokenRefreshURL, body)
}
