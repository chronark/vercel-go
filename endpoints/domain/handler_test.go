package domain_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/chronark/vercel-go/api"
	"github.com/chronark/vercel-go/utils/mocks"
	"github.com/stretchr/testify/require"

	"github.com/chronark/vercel-go/endpoints/domain"
	"github.com/google/uuid"
)

/* Run all tests on a personal account an on a team account */
type testCase struct {
	name    string
	handler *domain.DomainHandler
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
	teamid := os.Getenv("VERCEL_TEAM_ID")
	if teamid == "" {
		panic("VERCEL_TEAM_ID must be defined")
	}

	tests = []testCase{
		{name: "personal", handler: domain.New(api.New(api.NewClientConfig{Token: token}), "")},
		{name: "team", handler: domain.New(api.New(api.NewClientConfig{Token: token}), teamid)},
	}
}

func TestGetDomains(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, err := tt.handler.List(domain.ListDomainsRequest{})

			require.NoError(t, err)
			require.True(t, res.Pagination.Count > 0)
		})
	}

}

func TestListDomainsWithResponse400(t *testing.T) {

	client := &mocks.MockHTTPClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			res := &http.Response{}
			res.StatusCode = 400
			res.Body = io.NopCloser(strings.NewReader(("Hello World")))
			return res, nil
		},
	}
	handler := domain.New(api.New(api.NewClientConfig{HTTPClient: client}), "")

	_, err := handler.List(domain.ListDomainsRequest{})

	require.Error(t, err)

}

func TestAddDomain(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := fmt.Sprintf("%s.com", uuid.NewString())
			defer func() {
				_, err := tt.handler.Delete(domain.DeleteDomainRequest{Name: name})
				require.NoError(t, err)
			}()

			res, err := tt.handler.Add(domain.AddDomainRequest{Name: name})

			require.NoError(t, err)
			require.Equal(t, name, res.Domain.Name)
		})
	}
}

func TestGetDomain(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := fmt.Sprintf("%s.com", uuid.NewString())

			// Create the domain and clean it up at the end
			defer func() {
				_, err := tt.handler.Delete(domain.DeleteDomainRequest{Name: name})
				require.NoError(t, err)
			}()
			_, err := tt.handler.Add(domain.AddDomainRequest{Name: name})
			if err != nil {
				t.Error(err)
			}

			res, err := tt.handler.Get(domain.GetDomainRequest{Name: name})

			require.NoError(t, err)
			require.Equal(t, name, res.Domain.Name)
		})
	}

}

func TestDeleteDomain(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := fmt.Sprintf("%s.com", uuid.NewString())

			// Create the domain
			_, err := tt.handler.Add(domain.AddDomainRequest{Name: name})
			require.NoError(t, err)

			res, err := tt.handler.Delete(domain.DeleteDomainRequest{Name: name})
			require.NoError(t, err)

			require.True(t, len(res.Uid) > 0)

			// Now the domain should no longer exist
			_, err = tt.handler.Get(domain.GetDomainRequest{Name: name})
			require.Error(t, err)
		})
	}

}

func TestCheckAvailability(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := fmt.Sprintf("%s.com", uuid.NewString())

			// A new domain should be available
			res, err := tt.handler.CheckAvailability(domain.CheckAvailabilityRequest{Name: name})

			require.NoError(t, err)
			require.True(t, res.Available)

			// Add the domain
			defer func() {
				_, err := tt.handler.Delete(domain.DeleteDomainRequest{Name: name})
				require.NoError(t, err)
			}()
			_, err = tt.handler.Add(domain.AddDomainRequest{Name: name})
			require.NoError(t, err)

			// Now it should not be available
			res, err = tt.handler.CheckAvailability(domain.CheckAvailabilityRequest{Name: name})

			require.NoError(t, err)
			require.False(t, res.Available)
		})
	}

}

func TestCheckPrice(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name := fmt.Sprintf("%s.com", uuid.NewString())

			// A new domain should be available
			res, err := tt.handler.CheckPrice(domain.CheckPriceRequest{Name: name})

			require.NoError(t, err)
			require.True(t, res.Period > 0)
			require.True(t, res.Price > 0)
		})
	}

}
