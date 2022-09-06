package oauth2

import (
	"context"
	"errors"
	"github.com/joyqi/go-feishu/api"
	"sync"
	"time"
)

const TenantTokenURL = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

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

// TenantTokenSource returns a token source that retrieves tokens from the tenant token endpoint
func (c *Config) TenantTokenSource(ctx context.Context) TokenSource {
	c.once.Do(func() {
		c.tts = &tenantTokenSource{
			ctx:  ctx,
			conf: c,
		}
	})

	return c.tts
}

// tenantTokenSource is the token source of the tenant token endpoint
type tenantTokenSource struct {
	ctx  context.Context
	conf *Config
	t    *Token
	mu   sync.Mutex
}

// Token retrieves a token from the tenant token endpoint
func (s *tenantTokenSource) Token() (*Token, error) {
	defer s.mu.Unlock()

	s.mu.Lock()
	if s.t == nil || !s.t.Valid() {
		req := &TenantTokenRequest{
			AppID:     s.conf.AppID,
			AppSecret: s.conf.AppSecret,
		}

		resp := TenantTokenResponse{}
		err := httpPost(s.ctx, TenantTokenURL, req, &resp)
		if err != nil {
			return nil, err
		}

		if resp.Code != 0 {
			return nil, errors.New(resp.Msg)
		}

		s.t = &Token{
			AccessToken: resp.TenantAccessToken,
			Expiry:      time.Now().Add(time.Duration(resp.Expire) * time.Second),
		}
	}

	return s.t, nil
}

// Client returns a client that uses the tenant token source
func (s *tenantTokenSource) Client() api.Client {
	return &tokenClient{
		ctx: s.ctx,
		ts:  s,
	}
}
