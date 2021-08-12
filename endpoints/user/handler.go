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
func (h *UserHandler) Get() (userResponse UserResponse, err error) {
	err = h.vercelClient.Call("GET", "/www/user", nil, &userResponse)
	if err != nil {
		return UserResponse{}, fmt.Errorf("Unable to fetch user from vercel: %w", err)
	}
	return userResponse, nil
}
