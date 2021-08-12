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

// Return the team info
func (h *TeamHandler) Get(req GetTeamRequest) (res GetTeamResponse, err error) {
	path := "/v1/teams"

	// Teams can be references by their id or slug.
	// If id the url looks like this: .../<id>
	// Or when using a slug: ...?slug=<slug>
	if req.Id != "" {
		path = fmt.Sprintf("%s/%s", path, req.Id)
	} else if req.Slug != "" {
		path = fmt.Sprintf("/v1/teams?slug=%s", req.Slug)
	} else {
		return GetTeamResponse{}, fmt.Errorf("Unable to fetch team: Either `Id` or `Slug` must be defined")
	}

	err = h.vercelClient.Call("GET", path, nil, &res)
	if err != nil {
		return GetTeamResponse{}, fmt.Errorf("Unable to fetch team from vercel: %w", err)
	}
	return res, nil
}
