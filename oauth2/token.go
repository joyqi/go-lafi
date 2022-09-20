package oauth2

import (
	"github.com/joyqi/go-feishu/api/authen"
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

func NewToken(tk *authen.AccessTokenData) *Token {
	return &Token{
		AccessToken:  tk.AccessToken,
		RefreshToken: tk.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(tk.ExpiresIn) * time.Second),
	}
}
