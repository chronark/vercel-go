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

func TestGetDomains(t *testing.T) {
	handler := domain.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))

	res, err := handler.List(domain.ListDomainsRequest{})

	require.NoError(t, err)
	require.True(t, res.Pagination.Count > 0)
}

func TestListDomainsWithoutToken(t *testing.T) {

	handler := domain.New(api.New(api.NewClientConfig{Token: ""}))

	_, err := handler.List(domain.ListDomainsRequest{})

	require.Error(t, err)
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
	handler := domain.New(api.New(api.NewClientConfig{HTTPClient: client}))

	_, err := handler.List(domain.ListDomainsRequest{})

	require.Error(t, err)

}

func TestAddDomain(t *testing.T) {
	name := fmt.Sprintf("%s.com", uuid.NewString())
	handler := domain.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))
	defer func() {
		_, err := handler.Remove(domain.RemoveDomainRequest{Name: name})
		require.NoError(t, err)
	}()

	res, err := handler.Add(domain.AddDomainRequest{Name: name})

	require.NoError(t, err)
	require.Equal(t, name, res.Domain.Name)

}

func TestGetDomain(t *testing.T) {
	name := fmt.Sprintf("%s.com", uuid.NewString())
	handler := domain.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))

	// Create the domain and clean it up at the end
	defer func() {
		_, err := handler.Remove(domain.RemoveDomainRequest{Name: name})
		require.NoError(t, err)
	}()
	_, err := handler.Add(domain.AddDomainRequest{Name: name})
	if err != nil {
		t.Error(err)
	}

	res, err := handler.Get(domain.GetDomainRequest{Name: name})

	require.NoError(t, err)
	require.Equal(t, name, res.Domain.Name)

}

func TestRemoveDomain(t *testing.T) {
	name := fmt.Sprintf("%s.com", uuid.NewString())
	handler := domain.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))

	// Create the domain
	_, err := handler.Add(domain.AddDomainRequest{Name: name})
	require.NoError(t, err)

	res, err := handler.Remove(domain.RemoveDomainRequest{Name: name})
	require.NoError(t, err)

	require.True(t, len(res.Uid) > 0)

	// Now the domain should no longer exist
	_, err = handler.Get(domain.GetDomainRequest{Name: name})
	require.Error(t, err)

}

func TestCheckAvailability(t *testing.T) {
	name := fmt.Sprintf("%s.com", uuid.NewString())
	handler := domain.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))

	// A new domain should be available
	res, err := handler.CheckAvailability(domain.CheckAvailabilityRequest{Name: name})

	require.NoError(t, err)
	require.True(t, res.Available)

	// Add the domain
	defer func() {
		_, err := handler.Remove(domain.RemoveDomainRequest{Name: name})
		require.NoError(t, err)
	}()
	_, err = handler.Add(domain.AddDomainRequest{Name: name})
	require.NoError(t, err)

	// Now it should not be available
	res, err = handler.CheckAvailability(domain.CheckAvailabilityRequest{Name: name})

	require.NoError(t, err)
	require.False(t, res.Available)

}

func TestCheckPrice(t *testing.T) {
	name := fmt.Sprintf("%s.com", uuid.NewString())
	handler := domain.New(api.New(api.NewClientConfig{Token: os.Getenv("VERCEL_TOKEN")}))

	// A new domain should be available
	res, err := handler.CheckPrice(domain.CheckPriceRequest{Name: name})

	require.NoError(t, err)
	require.True(t, res.Period > 0)
	require.True(t, res.Price > 0)

}
