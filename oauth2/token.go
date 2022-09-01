package oauth2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joyqi/go-oauth2-feishu/http"
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
	Groups       []string
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

type UserGroupsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		GroupList []string `json:"group_list,flow"`
		HasMore   bool     `json:"has_more"`
	} `json:"data"`
}

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

func (s *TokenSource) valid() bool {
	return time.Now().Add(time.Minute).Before(s.t.Expiry)
}

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
	err = http.Post(
		ctx,
		endpointURL,
		req,
		&resp,
		http.Header{Key: "Authorization", Value: tenantToken},
	)

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}

	groups, err := retrieveUserGroups(resp.Data.OpenId, tenantToken)
	if err != nil {
		return nil, err
	}

	token := &Token{
		AccessToken:  resp.Data.AccessToken,
		RefreshToken: resp.Data.RefreshToken,
		Expiry:       time.Unix(resp.Data.ExpiresIn, 0),
		Groups:       groups,
	}

	return token, nil
}

// retrieveUserGroups retrieves the user groups associated with the given user ID.
func retrieveUserGroups(openId string, tenantToken string) ([]string, error) {
	url := fmt.Sprintf(EndpointURL.UserGroupsApiURL+"?member_id=%s", openId)

	err := http.Get(url, http.Header{Key: "Authorization", Value: "Bearer " + tenantToken})
	if err != nil {
		return nil, err
	}

	resp := UserGroupsResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, errors.New(resp.Msg)
	}

	return resp.Data.GroupList, nil
}
