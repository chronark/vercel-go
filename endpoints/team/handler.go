package team

import (
	"fmt"
	"github.com/chronark/vercel-go/api"
)

type TeamHandler struct {
	vercelClient api.VercelClient
}

func New(vercelClient api.VercelClient) *TeamHandler {
	return &TeamHandler{
		vercelClient,
	}
}

// Return the authenticated user
func (h *TeamHandler) Get(req GetTeamRequest) (res GetTeamResponse, err error) {

	if req.Id != "" {
		err = h.vercelClient.Call("GET", fmt.Sprintf("/v1/teams/%s", req.Id), nil, &res)
	} else if req.Slug != "" {
		fmt.Println(req.Slug)
		err = h.vercelClient.Call("GET", fmt.Sprintf("/v1/teams?slug=%s", req.Slug), nil, &res)
	} else {
		return GetTeamResponse{}, fmt.Errorf("Unable to fetch team: Either `Id` or `Slug` must be defined")
	}

	if err != nil {
		return GetTeamResponse{}, fmt.Errorf("Unable to fetch team from vercel: %w", err)
	}
	return res, nil
}
