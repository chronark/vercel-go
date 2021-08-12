package secret

import (
	"fmt"

	"github.com/chronark/vercel-go/api"
)

type SecretHandler struct {
	vercelClient api.VercelClient
}

func New(vercelClient api.VercelClient) *SecretHandler {
	return &SecretHandler{
		vercelClient,
	}
}

// Retrieves the active now secrets for the authenticating user.
func (h *SecretHandler) ListSecrets(req ListSecretsRequest) (res ListSecretsResponse, err error) {

	err = h.vercelClient.Call("GET", "/v3/now/secrets", nil, &res)

	if err != nil {
		return ListSecretsResponse{}, fmt.Errorf("Unable to fetch secrets from vercel: %w", err)
	}
	return res, nil
}

// Get the information for a specific secret by passing either the secret id or name.
func (h *SecretHandler) GetSecret(req GetSecretRequest) (res GetSecretResponse, err error) {
	var path string
	if req.Id != "" {
		path = fmt.Sprintf("/v3/now/secrets/%s", req.Id)
	} else if req.Name != "" {
		path = fmt.Sprintf("/v3/now/secrets/%s", req.Name)
	} else {
		return GetSecretResponse{}, fmt.Errorf("Unable to fetch team: Either `Id` or `Slug` must be defined")
	}
	err = h.vercelClient.Call("GET", path, nil, &res)

	if err != nil {
		return GetSecretResponse{}, fmt.Errorf("Unable to fetch secrets from vercel: %w", err)
	}
	return res, nil
}

// Retrieves the active now secrets for the authenticating user.
func (h *SecretHandler) Create(req CreateSecretsRequest) (res CreateSecretsResponse, err error) {

	err = h.vercelClient.Call("POST", "/v2/now/secrets", req, &res)

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

	err = h.vercelClient.Call("PATCH", fmt.Sprintf("/v2/now/secrets/%s", req.Name), Payload{Name: req.NewName}, &res)

	if err != nil {
		return RenameSecretResponse{}, fmt.Errorf("Unable to rename secret: %w", err)
	}
	return res, nil
}

// This endpoint provides an opportunity to edit the name of a user's secret.
// The name has to be unique to that user's secrets.
func (h *SecretHandler) Delete(req DeleteSecretRequest) (res DeleteSecretResponse, err error) {

	err = h.vercelClient.Call("DELETE", fmt.Sprintf("/v2/now/secrets/%s", req.Name), nil, &res)

	if err != nil {
		return DeleteSecretResponse{}, fmt.Errorf("Unable to delete secret: %w", err)
	}
	return res, nil
}
