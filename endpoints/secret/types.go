package secret

type Secret struct {
	// The unique identifier of the secret. Always prepended with sec_.
	Uid string `json:"uid"`

	// The name of the secret. This is what you could use in your environment variables after a @.
	Name string `json:"name"`

	// The date when the secret was created (ISO 8601).
	Created string `json:"created"`

	// The date when the secret was created (timestamp).
	CreatedAt int `json:"createdAt"`
}

type ListSecretsRequest struct {
	// Maximum number of domains to list from a request.
	// Required: No
	Limit int `json:"limit,omitempty"`

	// Get domains created after this JavaScript timestamp.
	// Required: No
	Since int `json:"since,omitempty"`

	// Get domains created before this JavaScript timestamp.
	// Required: No
	Until int `json:"until,omitempty"`
}

type ListSecretsResponse struct {
	Secrets    []Secret `json:"secrets"`
	Pagination struct {
		Count int   `json:"count"`
		Next  int64 `json:"next"`
		Prev  int64 `json:"prev"`
	} `json:"pagination"`
}

// Either `Name` or `Id` must be provided
type GetSecretRequest struct {
	// The name of the secret.
	// Required: false
	Name string

	// The unique identifier of the secret.
	// Required: false
	Id string
}

type GetSecretResponse struct {
	// The unique identifier of the secret.
	Uid string `json:"uid"`

	// The name of the secret.
	Name string `json:"name"`

	// The team unique identifier to which the secret belongs to.
	TeamId string `json:"teamId"`

	// The unique identifier of the user who created the secret.
	UserId string `json:"userId"`

	// The date when the secret was created.
	Created string `json:"created"`

	// The date when the secret was created in milliseconds since the UNIX epoch.
	CreatedAt int `json:"createdAt"`
}

type CreateSecretsRequest struct {
	// The name of the secret (max 100 characters).
	// Required: Yes
	Name string `json:"name"`
	// The value of the new secret.
	// Required: Yes
	Value string `json:"value"`
}

type CreateSecretsResponse struct {
	// The unique identifier of the secret. Always prepended with sec_.
	Uid string `json:"uid"`

	// The name of the secret.
	Name string `json:"name"`

	// The date when the secret was created.
	Created string `json:"created"`

	// The unique identifier of the user who created the secret.
	UserId string `json:"userId"`

	// A map with the value of the secret.
	Value struct {
		// The type of structure used to save the secret value (always Buffer).
		Type string `json:"type"`

		// 	A list of numbers which could be used to recreate the secret value.
		data []int `son:"data"`
	} `json:"value"`
}

type RenameSecretRequest struct {
	Name string

	// The new name
	// Required: Yes
	NewName string
}

type RenameSecretResponse struct {
	// The unique identifier of the secret. Always prepended with sec_.
	Uid string `json:"uid"`

	// The new name of the secret.
	Name string `json:"name"`

	// The date when the secret was created.
	Created string `json:"created"`

	// The old name of the secret.
	OldName string `json:"oldName"`
}

type DeleteSecretRequest struct {
	// The name of the secret
	// Required: Yes
	Name string `json:"name"`
}

type DeleteSecretResponse struct {
	// The unique identifier of the secret. Always prepended with sec_.
	Uid string `json:"uid"`

	// The new name of the secret.
	Name string `json:"name"`

	// The date when the secret was created.
	Created int `json:"created"`
}
