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
	// Set the teamId to perform all actions on behalf of a team.
	// This means you need multiple instances to handle resources for different
	// teams or your personal account
	TeamId string
}

func New(config NewClientConfig) *Client {
	api := api.New(api.NewClientConfig{Token: config.Token, UserAgent: config.UserAgent})
	return &Client{
		Domain: domain.New(api, config.TeamId),
		Secret: secret.New(api, config.TeamId),
		Team:   team.New(api),
		User:   user.New(api),
	}
}
