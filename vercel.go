package vercel

import (
	"github.com/chronark/vercel-go/api"
	"github.com/chronark/vercel-go/endpoints/domain"
	"github.com/chronark/vercel-go/endpoints/secret"
	"github.com/chronark/vercel-go/endpoints/team"
	"github.com/chronark/vercel-go/endpoints/user"
)

type Client struct {
	Domain *domain.DomainHandler
	Secret *secret.SecretHandler
	Team   *team.TeamHandler
	User   *user.UserHandler
}

type NewClientConfig struct {
	Token     string
	UserAgent string
}

func New(config NewClientConfig) *Client {
	api := api.New(api.NewClientConfig{Token: config.Token, UserAgent: config.UserAgent})
	return &Client{
		Domain: domain.New(api),
		Secret: secret.New(api),
		Team:   team.New(api),
		User:   user.New(api),
	}
}
