package oauth2

import (
	"context"
	"github.com/joyqi/go-feishu/api"
	"github.com/joyqi/go-feishu/api/authen"
	"github.com/joyqi/go-feishu/httptool"
	"net/url"
	"sync"
)

const AuthURL = "https://open.feishu.cn/open-apis/authen/v1/index"

// Config represents the configuration of the oauth2 service
type Config struct {
	// AppID is the app id of oauth2.
	AppID string

	// AppSecret is the app secret of oauth2.
	AppSecret string

	// AppTicket represents the ticket of the app.
	// It's used to retrieve the tenant access token.
	// If you're using the internal app, please leave it empty.
	AppTicket string

	// TenantKey represents the key of the tenant.
	TenantKey string

	// RedirectURL is the URL to redirect users going through
	RedirectURL string

	// once is used to ensure tts is initialized only once.
	once sync.Once

	// tts is the tenant token source
	tts ClientSource
}

// A TokenSource is anything that can return a token.
type TokenSource interface {
	Token() (*Token, error)
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
	req := &authen.AccessTokenCreateBody{
		GrantType: "authorization_code",
		Code:      code,
	}

	tokenApi := &authen.AccessToken{Client: c.TenantTokenSource(ctx).Client()}
	tk, err := tokenApi.Create(req)

	if err != nil {
		return nil, err
	}

	return NewToken(tk), nil
}

// TokenSource returns a TokenSource to grant token access
func (c *Config) TokenSource(ctx context.Context, t *Token) ClientSource {
	return &reuseTokenSource{
		ctx:  ctx,
		conf: c,
		t:    t,
	}
}

// reuseTokenSource is a TokenSource that reuse the token if it's still valid
type reuseTokenSource struct {
	ctx  context.Context
	conf *Config
	t    *Token
}

// Token returns a token if it's still valid, else will refresh the token
func (s *reuseTokenSource) Token() (*Token, error) {
	if !s.t.Valid() {
		token, err := s.refresh()
		if err != nil {
			return nil, err
		}

		s.t = token
	}

	return s.t, nil
}

// Client returns a client that uses the token source to retrieve the token
func (s *reuseTokenSource) Client() api.Client {
	return &tokenClient{
		ctx: s.ctx,
		ts:  s,
	}
}

// refresh retrieves the token from the endpoint
func (s *reuseTokenSource) refresh() (*Token, error) {
	req := &authen.AccessTokenRefreshBody{
		RefreshToken: s.t.RefreshToken,
		GrantType:    "refresh_token",
	}

	tokenApi := &authen.AccessToken{Client: s.conf.TenantTokenSource(s.ctx).Client()}
	tk, err := tokenApi.Refresh(req)

	if err != nil {
		return nil, err
	}

	return NewToken(tk), nil
}
