package auth

import (
	"encoding/json"
	"errors"
	"github.com/joyqi/go-feishu/api"
)

// TokenResponse represents the common token response structure
type TokenResponse struct {
	// Code is the response status code
	Code int `json:"code"`

	// Msg is the response message
	Msg string `json:"msg"`

	// AccessToken is the access token
	AccessToken string

	// Expire is the expiration time of the access token
	Expire int64 `json:"expire"`
}

// MakeTokenApi creates a new token api
func MakeTokenApi(c api.Client, tokenName string, uri string, body interface{}) (string, int64, error) {
	var resp map[string]json.RawMessage
	token := TokenResponse{}
	err := c.Request("POST", uri, body, &resp)

	if err = json.Unmarshal(resp["code"], &token.Code); err != nil {
		return "", 0, err
	}

	if err = json.Unmarshal(resp["msg"], &token.Msg); err != nil {
		return "", 0, err
	}

	if err = json.Unmarshal(resp[tokenName], &token.AccessToken); err != nil {
		return "", 0, err
	}

	if err = json.Unmarshal(resp["expire"], &token.Expire); err != nil {
		return "", 0, err
	}

	if token.Code != 0 {
		return "", 0, errors.New(token.Msg)
	}

	return token.AccessToken, token.Expire, nil
}
