package oauth2

import (
	"context"
	"errors"
	"time"
)

const TenantTokenURL = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

type TenantToken struct {
	AccessToken string
	Expiry      time.Time
}

// TenantTokenRequest represents a request to retrieve a tenant token
type TenantTokenRequest struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

// TenantTokenResponse is the token response of the tenant endpoint
type TenantTokenResponse struct {
	// Code is the response status code
	Code int `json:"code"`

	// Msg is the response message
	Msg string `json:"msg"`

	// TenantAccessToken is the access token
	TenantAccessToken string `json:"tenant_access_token"`

	// Expire is the expiration time of the access token
	Expire int64 `json:"expire"`
}

// TenantToken represents a tenant access token from tenant token endpoint
func (c *Config) TenantToken(ctx context.Context) (string, error) {
	defer c.tenantTokenMu.Unlock()

	c.tenantTokenMu.Lock()
	if !c.TenantTokenValid() {
		req := &TenantTokenRequest{
			AppID:     c.AppID,
			AppSecret: c.AppSecret,
		}

		resp := TenantTokenResponse{}
		err := httpPost(ctx, TenantTokenURL, req, &resp)
		if err != nil {
			return "", err
		}

		if resp.Code != 0 {
			return "", errors.New(resp.Msg)
		}

		c.tenantToken = resp.TenantAccessToken
		c.tenantTokenExpireAt = time.Now().Add(time.Duration(resp.Expire) * time.Second)
	}

	return c.tenantToken, nil
}

// ClientToken is an alias of TenantToken to implement the ClientTokenSource interface
func (c *Config) ClientToken(ctx context.Context) (string, error) {
	return c.TenantToken(ctx)
}

// TenantTokenValid returns true if the tenant token is valid
func (c *Config) TenantTokenValid() bool {
	return c.tenantToken != "" && time.Now().Add(time.Minute).Before(c.tenantTokenExpireAt)
}
