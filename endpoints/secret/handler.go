package secret

import (
	"fmt"

	"github.com/chronark/vercel-go/api"
)

type SecretHandler struct {
	vercelClient api.VercelClient
	teamid       string
}

func New(vercelClient api.VercelClient, teamid string) *SecretHandler {
	return &SecretHandler{
		vercelClient,
		teamid,
	}
}

// Retrieves the active now secrets for the authenticating user.
func (h *SecretHandler) ListSecrets(req ListSecretsRequest) (res ListSecretsResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v3/now/secrets", &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return ListSecretsResponse{}, fmt.Errorf("Unable to fetch secrets from vercel: %w", err)
	}
	return res, nil
}

// Get the information for a specific secret by passing either the secret id or name.
func (h *SecretHandler) GetSecret(secretIdOrName string) (res GetSecretResponse, err error) {

	path := fmt.Sprintf("/v3/now/secrets/%s", secretIdOrName)

	apiRequest := api.NewApiRequest("GET", path, &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}

	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return GetSecretResponse{}, fmt.Errorf("Unable to fetch secrets from vercel: %w", err)
	}
	return res, nil
}

// Retrieves the active now secrets for the authenticating user.
func (h *SecretHandler) Create(req CreateSecretsRequest) (res CreateSecretsResponse, err error) {

	apiRequest := api.NewApiRequest("POST", "/v2/now/secrets", &res)
	apiRequest.Body = req
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}

	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return CreateSecretsResponse{}, fmt.Errorf("Unable to create secret: %w", err)
	}
	return res, nil
}

// This endpoint provides an opportunity to edit the name of a user's secret.
// The name has to be unique to that user's secrets.
func (h *SecretHandler) Rename(req RenameSecretRequest) (res RenameSecretResponse, err error) {
	type Payload struct {
		Name string `json:"name"`
	}
	apiRequest := api.NewApiRequest("PATCH", fmt.Sprintf("/v2/now/secrets/%s", req.Name), &res)
	apiRequest.Body = Payload{Name: req.NewName}
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}

	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return RenameSecretResponse{}, fmt.Errorf("Unable to rename secret: %w", err)
	}
	return res, nil
}

// This endpoint provides an opportunity to edit the name of a user's secret.
// The name has to be unique to that user's secrets.
func (h *SecretHandler) Delete(req DeleteSecretRequest) (res DeleteSecretResponse, err error) {

	apiRequest := api.NewApiRequest("DELETE", fmt.Sprintf("/v2/now/secrets/%s", req.Name), &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return DeleteSecretResponse{}, fmt.Errorf("Unable to delete secret: %w", err)
	}
	return res, nil
}
