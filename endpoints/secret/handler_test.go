package secret_test

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/chronark/vercel-go/api"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/chronark/vercel-go/endpoints/secret"
	"github.com/chronark/vercel-go/utils/mocks"
)

/* Run all tests on a personal account an on a team account */
type testCase struct {
	name    string
	handler *secret.SecretHandler
}

var (
	token string
	tests []testCase
)

func init() {
	token = os.Getenv("VERCEL_TOKEN")
	if token == "" {
		panic("VERCEL_TOKEN must be defined")
	}
	teamId := os.Getenv("VERCEL_TEAM_ID")
	if teamId == "" {
		panic("VERCEL_TEAM_ID must be defined")
	}

	tests = []testCase{
		{name: "personal", handler: secret.New(api.New(api.NewClientConfig{Token: token}), "")},
		{name: "team", handler: secret.New(api.New(api.NewClientConfig{Token: token}), teamId)},
	}
}

func TestListSecretsWithResponse400(t *testing.T) {
	client := &mocks.MockHTTPClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			res := &http.Response{}
			res.StatusCode = 400
			res.Body = io.NopCloser(strings.NewReader(("Hello World")))
			return res, nil
		},
	}
	handler := secret.New(api.New(api.NewClientConfig{HTTPClient: client}), "")

	_, err := handler.ListSecrets(secret.ListSecretsRequest{})

	require.Error(t, err)

}
func TestListSecrets(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := uuid.NewString()

			//Create at least 1 secret and delete it afterwards
			defer func() {
				_, err := tt.handler.Delete(secret.DeleteSecretRequest{Name: name})
				require.NoError(t, err)
			}()
			_, err := tt.handler.Create(secret.CreateSecretsRequest{Name: name, Value: "secret"})
			require.NoError(t, err)

			res, err := tt.handler.ListSecrets(secret.ListSecretsRequest{})

			require.NoError(t, err)
			require.True(t, len(res.Secrets) > 0)
		})
	}
}

func TestGetSecretWithName(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := uuid.NewString()

			//Create at least 1 secret and delete it afterwards
			defer func() {
				_, err := tt.handler.Delete(secret.DeleteSecretRequest{Name: name})
				require.NoError(t, err)
			}()
			_, err := tt.handler.Create(secret.CreateSecretsRequest{Name: name, Value: "secret"})
			require.NoError(t, err)

			res, err := tt.handler.GetSecret(secret.GetSecretRequest{Name: name})

			require.NoError(t, err)
			require.True(t, len(res.Name) > 0)
		})
	}
}

func TestGetSecretWithId(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := uuid.NewString()

			//Create at least 1 secret and delete it afterwards
			defer func() {
				_, err := tt.handler.Delete(secret.DeleteSecretRequest{Name: name})
				require.NoError(t, err)
			}()
			createRes, err := tt.handler.Create(secret.CreateSecretsRequest{Name: name, Value: "secret"})
			require.NoError(t, err)

			res, err := tt.handler.GetSecret(secret.GetSecretRequest{Id: createRes.Uid})

			require.NoError(t, err)
			require.True(t, len(res.Name) > 0)
		})
	}
}

func TestCreateSecret(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			name := uuid.NewString()
			defer func() {
				_, err := tt.handler.Delete(secret.DeleteSecretRequest{Name: name})
				require.NoError(t, err)
			}()
			res, err := tt.handler.Create(secret.CreateSecretsRequest{Name: name, Value: "value"})

			require.NoError(t, err)
			foundRes, err := tt.handler.GetSecret(secret.GetSecretRequest{Name: name})
			require.NoError(t, err)
			require.Equal(t, res.Uid, foundRes.Uid)
		})
	}
}

func TestRenameSecret(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := uuid.NewString()
			newName := uuid.NewString()
			defer func() {
				_, err := tt.handler.Delete(secret.DeleteSecretRequest{Name: newName})
				require.NoError(t, err)
			}()
			createRes, err := tt.handler.Create(secret.CreateSecretsRequest{Name: name, Value: "value"})
			require.NoError(t, err)

			_, err = tt.handler.Rename(secret.RenameSecretRequest{Name: name, NewName: newName})
			require.NoError(t, err)

			foundRes, err := tt.handler.GetSecret(secret.GetSecretRequest{Name: newName})
			require.NoError(t, err)
			require.Equal(t, createRes.Uid, foundRes.Uid)
		})
	}
}
