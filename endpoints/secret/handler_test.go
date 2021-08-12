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

func TestListSecretsWithResponse400(t *testing.T) {
	client := &mocks.MockHTTPClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			res := &http.Response{}
			res.StatusCode = 400
			res.Body = io.NopCloser(strings.NewReader(("Hello World")))
			return res, nil
		},
	}
	handler := secret.New(api.New(api.NewClientConfig{HTTPClient: client}))

	_, err := handler.ListSecrets(secret.ListSecretsRequest{})

	require.Error(t, err)

}
func TestListSecrets(t *testing.T) {
	handler := secret.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))
	name := uuid.NewString()

	//Create at least 1 secret and delete it afterwards
	defer func() {
		_, err := handler.Delete(secret.DeleteSecretRequest{Name: name})
		require.NoError(t, err)
	}()
	_, err := handler.Create(secret.CreateSecretsRequest{Name: name, Value: "secret"})
	require.NoError(t, err)

	res, err := handler.ListSecrets(secret.ListSecretsRequest{})

	require.NoError(t, err)
	require.True(t, len(res.Secrets) > 0)
}

func TestGetSecretWithName(t *testing.T) {
	handler := secret.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))
	name := uuid.NewString()

	//Create at least 1 secret and delete it afterwards
	defer func() {
		_, err := handler.Delete(secret.DeleteSecretRequest{Name: name})
		require.NoError(t, err)
	}()
	_, err := handler.Create(secret.CreateSecretsRequest{Name: name, Value: "secret"})
	require.NoError(t, err)

	res, err := handler.GetSecret(secret.GetSecretRequest{Name: name})

	require.NoError(t, err)
	require.True(t, len(res.Name) > 0)
}

func TestGetSecretWithId(t *testing.T) {
	handler := secret.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))
	name := uuid.NewString()

	//Create at least 1 secret and delete it afterwards
	defer func() {
		_, err := handler.Delete(secret.DeleteSecretRequest{Name: name})
		require.NoError(t, err)
	}()
	createRes, err := handler.Create(secret.CreateSecretsRequest{Name: name, Value: "secret"})
	require.NoError(t, err)

	res, err := handler.GetSecret(secret.GetSecretRequest{Id: createRes.Uid})

	require.NoError(t, err)
	require.True(t, len(res.Name) > 0)
}

func TestCreateSecret(t *testing.T) {
	handler := secret.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))
	name := uuid.NewString()
	defer func() {
		_, err := handler.Delete(secret.DeleteSecretRequest{Name: name})
		require.NoError(t, err)
	}()
	res, err := handler.Create(secret.CreateSecretsRequest{Name: name, Value: "value"})

	require.NoError(t, err)
	foundRes, err := handler.GetSecret(secret.GetSecretRequest{Name: name})
	require.NoError(t, err)
	require.Equal(t, res.Uid, foundRes.Uid)
}

func TestRenameSecret(t *testing.T) {
	handler := secret.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))
	name := uuid.NewString()
	newName := uuid.NewString()
	defer func() {
		_, err := handler.Delete(secret.DeleteSecretRequest{Name: newName})
		require.NoError(t, err)
	}()
	createRes, err := handler.Create(secret.CreateSecretsRequest{Name: name, Value: "value"})
	require.NoError(t, err)

	_, err = handler.Rename(secret.RenameSecretRequest{Name: name, NewName: newName})
	require.NoError(t, err)

	foundRes, err := handler.GetSecret(secret.GetSecretRequest{Name: newName})
	require.NoError(t, err)
	require.Equal(t, createRes.Uid, foundRes.Uid)

}
