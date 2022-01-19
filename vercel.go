package vercel

import (
	"github.com/chronark/vercel-go/api"
	"github.com/chronark/vercel-go/endpoints/dns"
	"github.com/chronark/vercel-go/endpoints/domain"
	"github.com/chronark/vercel-go/endpoints/project"
	"github.com/chronark/vercel-go/endpoints/secret"
	"github.com/chronark/vercel-go/endpoints/team"
	"github.com/chronark/vercel-go/endpoints/user"
)

type Client struct {
	Domain  *domain.DomainHandler
	Secret  *secret.SecretHandler
	Team    *team.TeamHandler
	User    *user.UserHandler
	Project *project.ProjectHandler
	Dns     *dns.DnsHandler
}

type NewClientConfig struct {
	Token     string
	UserAgent string
	// Set the teamid to perform all actions on behalf of a team.
	// This means you need multiple instances to handle resources for different
	// teams or your personal account
	Teamid string
}

func New(config NewClientConfig) *Client {
	api := api.New(api.NewClientConfig{Token: config.Token, UserAgent: config.UserAgent})
	return &Client{
		Domain:  domain.New(api, config.Teamid),
		Secret:  secret.New(api, config.Teamid),
		Team:    team.New(api),
		User:    user.New(api),
		Project: project.New(api, config.Teamid),
		Dns:     dns.New(api, config.Teamid),
	}
}
