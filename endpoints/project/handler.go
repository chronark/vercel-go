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
	apiRequest := api.NewApiRequest("POST", "/v8/projects", &res)
	apiRequest.Body = req
	if h.teamId != "" {
		apiRequest.Query.Add("teamId", h.teamId)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return CreateProjectResponse{}, fmt.Errorf("Unable to create project: %w", err)
	}
	return res, nil
}

// Get a list of the projects you currently have under your account.
// By default, it returns the last 20 projects. The rest can be retrieved
// using the pagination options described below.
// NOTE: The order is always based on the updatedAt field of the project.
func (h *ProjectHandler) ListProjects(req ListProjectsRequest) (res ListProjectsResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v8/projects", &res)
	if h.teamId != "" {
		apiRequest.Query.Add("teamId", h.teamId)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return ListProjectsResponse{}, fmt.Errorf("Unable to fetch projects: %w", err)
	}
	return res, nil
}

// Get the information for a specific project by passing either the project id
// or name in the URL.
func (h *ProjectHandler) GetProject(req GetProjectRequest) (res GetProjectResponse, err error) {

	path := "/v8/projects"

	if req.Id != "" {
		path = fmt.Sprintf("%s/%s", path, req.Id)
	} else if req.Name != "" {
		path = fmt.Sprintf("%s/%s", path, req.Name)
	} else {
		return GetProjectResponse{}, fmt.Errorf("Unable to load project: Either `Id` or `Name` must be defined")
	}

	apiRequest := api.NewApiRequest("GET", path, &res)
	if h.teamId != "" {
		apiRequest.Query.Add("teamId", h.teamId)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return GetProjectResponse{}, fmt.Errorf("Unable to create project: %w", err)
	}
	return res, nil
}
