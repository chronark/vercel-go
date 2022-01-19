package project

import (
	"fmt"

	"github.com/chronark/vercel-go/api"
)

type ProjectHandler struct {
	vercelClient api.VercelClient
	teamid       string
}

func New(vercelClient api.VercelClient, teamid string) *ProjectHandler {
	return &ProjectHandler{
		vercelClient,
		teamid,
	}
}

// Create a new project
func (h *ProjectHandler) Create(req CreateProjectRequest) (res CreateProjectResponse, err error) {
	apiRequest := api.NewApiRequest("POST", "/v8/projects", &res)
	apiRequest.Body = req
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
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
func (h *ProjectHandler) List(req ListProjectsRequest) (res ListProjectsResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v8/projects", &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return ListProjectsResponse{}, fmt.Errorf("Unable to fetch projects: %w", err)
	}
	return res, nil
}

// Get the information for a specific project by passing either the project id
// or name in the URL.
func (h *ProjectHandler) Get(projectIdOrName string) (res GetProjectResponse, err error) {

	path := fmt.Sprintf("/v8/projects/%s", projectIdOrName)

	apiRequest := api.NewApiRequest("GET", path, &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return GetProjectResponse{}, fmt.Errorf("Unable to create project: %w", err)
	}
	return res, nil
}

// Get the information for a specific project by passing either the project id
// or name in the URL.
func (h *ProjectHandler) Delete(projectIdOrName string) (res DeleteProjectResponse, err error) {

	path := fmt.Sprintf("/v8/projects/%s", projectIdOrName)

	apiRequest := api.NewApiRequest("DELETE", path, nil)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return DeleteProjectResponse{}, fmt.Errorf("Unable to delete project: %w", err)
	}
	return res, nil
}

// Update the fields of a project using either its name or id.
func (h *ProjectHandler) Update(projectIdOrName string, req UpdateProjectRequest) (res UpdateProjectResponse, err error) {

	path := fmt.Sprintf("/v8/projects/%s", projectIdOrName)

	apiRequest := api.NewApiRequest("PATCH", path, &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	apiRequest.Body = req
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return UpdateProjectResponse{}, fmt.Errorf("Unable to update project: %w", err)
	}
	return res, nil
}

// Update the fields of a project using either its name or id.
func (h *ProjectHandler) CreateEnvironmentVariable(projectIdOrName string, req CreateEnvironmentVariableRequest) (res CreateEnvironmentVariableResponse, err error) {

	path := fmt.Sprintf("/v8/projects/%s/env", projectIdOrName)

	apiRequest := api.NewApiRequest("POST", path, &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	apiRequest.Body = req
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return CreateEnvironmentVariableResponse{}, fmt.Errorf("Unable to create environment variable: %w", err)
	}
	return res, nil
}

// Update the fields of a project using either its name or id.
func (h *ProjectHandler) DeleteEnvironmentVariable(req DeleteEnvironmentVariableRequest) (res DeleteEnvironmentVariableResponse, err error) {

	path := fmt.Sprintf("/v8/projects/%s/env/%s", req.Projectid, req.Envid)

	apiRequest := api.NewApiRequest("DELETE", path, nil)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return DeleteEnvironmentVariableResponse{}, fmt.Errorf("Unable to delete environment variable: %w", err)
	}
	return res, nil
}

// Update the fields of a project using either its name or id.
func (h *ProjectHandler) UpdateEnvironmentVariable(projectNameOrId string, req UpdateEnvironmentVariableRequest) (res UpdateEnvironmentVariableResponse, err error) {

	path := fmt.Sprintf("/v8/projects/%s/env/%s", projectNameOrId, req.Envid)

	apiRequest := api.NewApiRequest("PATCH", path, &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	apiRequest.Body = req
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return UpdateEnvironmentVariableResponse{}, fmt.Errorf("Unable to update environment variable: %w", err)
	}
	return res, nil
}

// Update the fields of a project using either its name or id.
func (h *ProjectHandler) AddDomain(projectIdOrName string, req AddDomainRequest) (res Domain, err error) {

	path := fmt.Sprintf("/v8/projects/%s/domains", projectIdOrName)

	apiRequest := api.NewApiRequest("POST", path, &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	apiRequest.Body = req
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return Domain{}, fmt.Errorf("Unable to add domain: %w", err)
	}
	return res, nil
}

func (h *ProjectHandler) GetDomain(projectIdOrName string, domain string) (res Domain, err error) {

	path := fmt.Sprintf("/v8/projects/%s/domains/%s", projectIdOrName, domain)

	apiRequest := api.NewApiRequest("GET", path, &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return Domain{}, fmt.Errorf("Unable to get domain: %w", err)
	}
	return res, nil
}

// Update the fields of a project using either its name or id.
func (h *ProjectHandler) UpdateDomain(projectNameOrId string, req UpdateDomainRequest) (res Domain, err error) {

	path := fmt.Sprintf("/v8/projects/%s/domains/%s", projectNameOrId, req.Name)

	apiRequest := api.NewApiRequest("PATCH", path, &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	apiRequest.Body = req
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return Domain{}, fmt.Errorf("Unable to update domain: %w", err)
	}
	return res, nil
}

func (h *ProjectHandler) ListDomains(projectIdOrName string, req ListDomainsRequest) (res ListDomainsResponse, err error) {
	apiRequest := api.NewApiRequest("GET", fmt.Sprintf("/v8/projects/%s/domains", projectIdOrName), &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return ListDomainsResponse{}, fmt.Errorf("Unable to fetch projects: %w", err)
	}
	return res, nil
}
