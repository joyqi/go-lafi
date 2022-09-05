package oauth2

import (
	"context"
	"github.com/joyqi/go-feishu/httptool"
	"net/url"
	"sync"
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

	// once is used to ensure tts is initialized only once.
	once sync.Once

	// tts is the tenant token source
	tts TokenSource
}

// A TokenSource is anything that can return a token.
type TokenSource interface {
	Token() (*Token, error)
}

// TokenRequest represents a request to retrieve a token from the server
type TokenRequest struct {
	GrantType string `json:"grant_type"`
	Code      string `json:"code"`
}

// RefreshTokenRequest represents a request to refresh the token from the server
type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

// TokenResponse represents the response from the Token service
type TokenResponse struct {
	// Code is the response status code
	Code int `json:"code"`

	// Msg is the response message in the response body
	Msg string `json:"msg"`

	// Data is the response body data
	Data struct {
		// OpenId represents the open ID of the user
		OpenId string `json:"open_id"`

		// AccessToken is the token used to access the application
		AccessToken string `json:"access_token"`

		// RefreshToken is the token used to refresh the user's access token
		RefreshToken string `json:"refresh_token"`

		// ExpiresIn is the number of seconds the token will be valid
		ExpiresIn int64 `json:"expires_in"`
	} `json:"data"`
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

	return retrieveToken(ctx, TokenURL, req, c.TenantTokenSource(ctx))
}

// TokenSource returns a TokenSource to grant token access
func (c *Config) TokenSource(ctx context.Context, t *Token) TokenSource {
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

// refresh retrieves the token from the endpoint
func (s *reuseTokenSource) refresh() (*Token, error) {
	req := &RefreshTokenRequest{
		RefreshToken: s.t.RefreshToken,
		GrantType:    "refresh_token",
	}

	return retrieveToken(s.ctx, RefreshTokenURL, req, s.conf.TenantTokenSource(s.ctx))
}
