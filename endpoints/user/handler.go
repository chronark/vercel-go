package user

import (
	"encoding/json"
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
func (h *UserHandler) Get() (user User, err error) {
	res, err := h.vercelClient.Call("GET", "/www/user", nil)
	if err != nil {
		return User{}, fmt.Errorf("Unable to fetch user from vercel: %w", err)
	}
	defer res.Body.Close()

	type UserResponse struct {
		User User `json:"user"`
	}

	var userResponse UserResponse
	err = json.NewDecoder(res.Body).Decode(&userResponse)
	if err != nil {
		return User{}, fmt.Errorf("Unable to unmarshal project: %w", err)
	}
	return userResponse.User, nil
}
