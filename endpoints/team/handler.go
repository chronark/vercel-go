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
	apiRequest := api.NewApiRequest("GET", "/v1/teams", &res)

	// Teams can be references by their id or slug.
	// If id the url looks like this: .../<id>
	// Or when using a slug: ...?slug=<slug>
	if req.Id != "" {
		apiRequest.Path = fmt.Sprintf("%s/%s", apiRequest.Path, req.Id)
	} else if req.Slug != "" {

		apiRequest.Query.Add("slug", req.Slug)
	} else {
		return GetTeamResponse{}, fmt.Errorf("Unable to fetch team: Either `Id` or `Slug` must be defined")
	}

	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return GetTeamResponse{}, fmt.Errorf("Unable to fetch team from vercel: %w", err)
	}
	return res, nil
}
