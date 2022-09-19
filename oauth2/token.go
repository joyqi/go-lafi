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

// Valid checks if the token is still valid
func (t *Token) Valid() bool {
	return time.Now().Add(time.Minute).Before(t.Expiry)
}

// retrieveToken retrieves the token from the endpoint
func retrieveToken(ctx context.Context, endpointURL string, req interface{}, ts TokenSource) (*Token, error) {
	t, err := ts.Token()
	if err != nil {
		return nil, err
	}

	resp := TokenResponse{}
	err = httpPost(
		ctx,
		endpointURL,
		req,
		&resp,
		httptool.Header{Key: "Authorization", Value: "Bearer " + t.AccessToken},
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
