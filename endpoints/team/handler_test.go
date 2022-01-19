package team_test

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/chronark/vercel-go/api"
	"github.com/chronark/vercel-go/utils/mocks"
	"github.com/stretchr/testify/require"

	"github.com/chronark/vercel-go/endpoints/team"
)

func TestGetTeamWithSlug(t *testing.T) {
	handler := team.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))

	res, err := handler.Get(team.GetTeamRequest{Slug: "perfolio"})

	require.NoError(t, err)
	require.True(t, len(res.Name) > 0)
}

func TestGetTeamWithid(t *testing.T) {
	handler := team.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))

	res, err := handler.Get(team.GetTeamRequest{Id: "team_ZYPNERZjT0L9LkPdP0Y9ZtS2"})

	require.NoError(t, err)
	require.True(t, len(res.Name) > 0)
}

func TestGetTeamWithoutidOrSlug(t *testing.T) {
	handler := team.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))

	_, err := handler.Get(team.GetTeamRequest{})

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
	handler := team.New(api.New(api.NewClientConfig{HTTPClient: client}))

	_, err := handler.Get(team.GetTeamRequest{Slug: "slug"})

	require.Error(t, err)

}
