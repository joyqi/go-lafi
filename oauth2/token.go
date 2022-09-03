package oauth2

import (
	"context"
	"errors"
	"github.com/joyqi/go-feishu/httptool"
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
}

type TokenSource struct {
	ctx  context.Context
	conf *Config
	t    *Token
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

// Token returns a token if it's still valid, else will refresh the token
func (s *TokenSource) Token() (*Token, error) {
	if !s.valid() {
		token, err := s.refresh()
		if err != nil {
			return nil, err
		}

		s.t = token
	}

	return s.t, nil
}

// ClientToken is an alias of Token to implement the ClientTokenSource interface
func (s *TokenSource) ClientToken(ctx context.Context) (string, error) {
	tk, err := s.Token()
	if err != nil {
		return "", err
	}

	return tk.AccessToken, nil
}

// valid checks if the token is still valid
func (s *TokenSource) valid() bool {
	return time.Now().Add(time.Minute).Before(s.t.Expiry)
}

// refresh retrieves the token from the endpoint
func (s *TokenSource) refresh() (*Token, error) {
	req := &RefreshTokenRequest{
		RefreshToken: s.t.RefreshToken,
		GrantType:    "refresh_token",
	}

	return retrieveToken(s.ctx, EndpointURL.RefreshTokenURL, req, s.conf)
}

// retrieveToken retrieves the token from the endpoint
func retrieveToken(ctx context.Context, endpointURL string, req interface{}, conf *Config) (*Token, error) {
	tenantToken, err := conf.TenantToken(ctx)
	if err != nil {
		return nil, err
	}

	resp := TokenResponse{}
	err = httpPost(
		ctx,
		endpointURL,
		req,
		&resp,
		httptool.Header{Key: "Authorization", Value: tenantToken},
	)

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}

	token := &Token{
		AccessToken:  resp.Data.AccessToken,
		RefreshToken: resp.Data.RefreshToken,
		Expiry:       time.Unix(resp.Data.ExpiresIn, 0),
	}

	return token, nil
}
