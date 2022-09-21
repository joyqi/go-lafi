package oauth2

import (
	"context"
	"github.com/joyqi/go-feishu/api"
	"github.com/joyqi/go-feishu/api/auth"
	"sync"
	"time"
)

// TenantTokenSource returns a token source that retrieves tokens from the tenant token endpoint
func (c *Config) TenantTokenSource(ctx context.Context) ClientSource {
	c.once.Do(func() {
		var ats *appTokenSource

		if c.AppTicket != "" && c.TenantKey != "" {
			ats = &appTokenSource{
				ctx:  ctx,
				conf: c,
			}
		}

		c.tts = &tenantTokenSource{
			ctx:  ctx,
			conf: c,
			ats:  ats,
		}
	})

	return c.tts
}

// tenantTokenSource is the token source of the tenant token endpoint
type tenantTokenSource struct {
	ctx  context.Context
	conf *Config
	t    *Token
	ats  *appTokenSource
	mu   sync.Mutex
}

// Token retrieves a token from the tenant token endpoint
func (s *tenantTokenSource) Token() (*Token, error) {
	defer s.mu.Unlock()

	s.mu.Lock()
	if s.t == nil || !s.t.Valid() {
		var (
			err error
			t   *Token
		)

		if s.ats != nil {
			t, err = s.retrieveCommonToken()
		} else {
			t, err = s.retrieveInternalToken()
		}

		if err != nil {
			return nil, err
		}

		s.t = t
	}

	return s.t, nil
}

// Client returns a client that uses the tenant token source
func (s *tenantTokenSource) Client() api.Client {
	return &tokenClient{
		ctx: s.ctx,
		ts:  s,
		t:   s.conf.Type,
	}
}

func (s *tenantTokenSource) retrieveInternalToken() (*Token, error) {
	c := &simpleClient{ctx: s.ctx}
	authApi := &auth.Tenant{Client: c}

	tk, expire, err := authApi.InternalAccessToken(&auth.TenantInternalBody{
		AppID:     s.conf.AppID,
		AppSecret: s.conf.AppSecret,
	})

	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: tk,
		Expiry:      time.Now().Add(time.Duration(expire) * time.Second),
	}, nil
}

func (s *tenantTokenSource) retrieveCommonToken() (*Token, error) {
	t, err := s.ats.Token()
	if err != nil {
		return nil, err
	}

	c := &simpleClient{ctx: s.ctx}
	authApi := &auth.Tenant{Client: c}

	tk, expire, err := authApi.CommonAccessToken(&auth.TenantCommonBody{
		AppAccessToken: t.AccessToken,
		TenantKey:      s.conf.TenantKey,
	})

	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: tk,
		Expiry:      time.Now().Add(time.Duration(expire) * time.Second),
	}, nil
}

type appTokenSource struct {
	ctx  context.Context
	conf *Config
	t    *Token
}

// Token retrieves a token from the app token endpoint
func (s *appTokenSource) Token() (*Token, error) {
	if s.t == nil || !s.t.Valid() {
		c := &simpleClient{ctx: s.ctx}
		authApi := &auth.App{Client: c}

		tk, expire, err := authApi.CommonAccessToken(&auth.AppCommonBody{
			AppID:     s.conf.AppID,
			AppSecret: s.conf.AppSecret,
			AppTicket: s.conf.AppTicket,
		})

		if err != nil {
			return nil, err
		}

		s.t = &Token{
			AccessToken: tk,
			Expiry:      time.Now().Add(time.Duration(expire) * time.Second),
		}
	}

	return s.t, nil
}
