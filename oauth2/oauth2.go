package oauth2

import (
	"context"
	"github.com/joyqi/go-feishu/httptool"
	"net/url"
	"sync"
	"time"
)

const (
	AuthURL         = "https://open.feishu.cn/open-apis/authen/v1/index"
	TokenURL        = "https://open.feishu.cn/open-apis/authen/v1/access_token"
	RefreshTokenURL = "https://open.feishu.cn/open-apis/authen/v1/refresh_access_token"
)

// Config represents the configuration of the oauth2 service
type Config struct {
	// AppID is the app id of oauth2.
	AppID string

	// AppSecret is the app secret of oauth2.
	AppSecret string

	// RedirectURL is the URL to redirect users going through
	RedirectURL string

	// tenantTokenMu is the lock for tenant token request
	tenantTokenMu sync.Mutex

	// tenantToken is the tenant token
	tenantToken string

	// tenantTokenExpireAt is the expiration time of the tenant token
	tenantTokenExpireAt time.Time
}

// AuthCodeURL is the URL to redirect users going through authentication
func (c *Config) AuthCodeURL(state string) string {
	v := url.Values{
		"response_type": {"code"},
		"app_id":        {c.AppID},
	}
	if c.RedirectURL != "" {
		v.Set("redirect_uri", c.RedirectURL)
	}

	if state != "" {
		v.Set("state", state)
	}

	return httptool.MakeURL(AuthURL, v)
}

// Exchange retrieve the token from access token endpoint
func (c *Config) Exchange(ctx context.Context, code string) (*Token, error) {
	req := &TokenRequest{
		GrantType: "authorization_code",
		Code:      code,
	}

	return retrieveToken(ctx, TokenURL, req, c)
}

// TokenSource returns a TokenSource to grant token access
func (c *Config) TokenSource(ctx context.Context, t *Token) *TokenSource {
	return &TokenSource{
		ctx:  ctx,
		conf: c,
		t:    t,
	}
}
