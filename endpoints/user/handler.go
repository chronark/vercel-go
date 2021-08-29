package user

import (
	"fmt"
	"github.com/chronark/vercel-go/api"
)

type UserHandler struct {
	vercelClient api.VercelClient
}

func New(vercelClient api.VercelClient) *UserHandler {
	return &UserHandler{
		vercelClient,
	}
}

// Return the authenticated user
func (h *UserHandler) Get() (res UserResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/www/user", &res)
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return UserResponse{}, fmt.Errorf("Unable to fetch user from vercel: %w", err)
	}
	return res, nil
}
