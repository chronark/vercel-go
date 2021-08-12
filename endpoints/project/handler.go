package project

import (
	"fmt"
	"github.com/chronark/vercel-go/api"
)

type ProjectHandler struct {
	vercelClient api.VercelClient
	teamId       string
}

func New(vercelClient api.VercelClient, teamId string) *ProjectHandler {
	return &ProjectHandler{
		vercelClient,
		teamId,
	}
}

// Create a new project
func (h *ProjectHandler) Create(req CreateProjectRequest) (res CreateProjectResponse, err error) {
	path := "/v8/projects"
	if h.teamId != "" {
		path = fmt.Sprintf("%s?teamId=%s", path, h.teamId)
	}
	err = h.vercelClient.Call("POST", path, req, &res)
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("Unable to create project: %w", err)
	}
	return res, nil
}
