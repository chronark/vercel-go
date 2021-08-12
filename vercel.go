package vercel

import (
	"github.com/chronark/vercel-go/endpoints/user"
	"github.com/chronark/vercel-go/api"
)

type Client struct {
	User *user.UserHandler
}

type NewClientConfig struct {
	Token     string
	UserAgent string
}

func New(config NewClientConfig) *Client {
	api := api.New(api.NewClientConfig{Token: config.Token, UserAgent: config.UserAgent})
	return &Client{
		User: user.New(api),
	}
}
