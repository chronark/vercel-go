package user_test

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/chronark/vercel-go/api"
	"github.com/chronark/vercel-go/utils/mocks"
	"github.com/stretchr/testify/require"

	"github.com/chronark/vercel-go/endpoints/user"
)

func TestGetUser(t *testing.T) {
	token := os.Getenv("VERCEL_TOKEN")
	if token == "" {
		t.Error("VERCEL_TOKEN must be defined")
	}
	handler := user.New(api.New(api.NewClientConfig{Token: token}))

	user, err := handler.Get()

	require.NoError(t, err)
	require.Equal(t, "chronark", user.Username)
}

func TestGetUserWithoutToken(t *testing.T) {

	handler := user.New(api.New(api.NewClientConfig{Token: ""}))

	_, err := handler.Get()

	require.Error(t, err)
}

func TestGetUserWithWrongToken(t *testing.T) {

	handler := user.New(api.New(api.NewClientConfig{Token: "wrongToken"}))

	_, err := handler.Get()

	require.Error(t, err)
}

func TestGetUserWithResponse400(t *testing.T) {
	client := &mocks.MockHTTPClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			res := &http.Response{}
			res.StatusCode = 400
			res.Body = io.NopCloser(strings.NewReader(("Hello World")))
			return res, nil
		},
	}
	handler := user.New(api.New(api.NewClientConfig{HTTPClient: client}))

	_, err := handler.Get()

	require.Error(t, err)
}
