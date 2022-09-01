package api

import "github.com/joyqi/go-oauth2-feishu/api/contact"

type Api interface {
	contact.Group
}

type Client struct {
}