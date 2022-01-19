package dns_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/chronark/vercel-go/api"
	"github.com/stretchr/testify/require"

	"github.com/chronark/vercel-go/endpoints/dns"
	"github.com/chronark/vercel-go/endpoints/domain"
	"github.com/google/uuid"
)

/* Run all tests on a personal account an on a team account */
type testCase struct {
	name          string
	dnsHandler    *dns.DnsHandler
	domainHandler *domain.DomainHandler
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
		{
			name:          "personal",
			domainHandler: domain.New(api.New(api.NewClientConfig{Token: token}), ""),
			dnsHandler:    dns.New(api.New(api.NewClientConfig{Token: token}), ""),
		},
		{
			name:          "team",
			domainHandler: domain.New(api.New(api.NewClientConfig{Token: token}), teamid),
			dnsHandler:    dns.New(api.New(api.NewClientConfig{Token: token}), teamid),
		},
	}

}

func TestCreateDnsRecord(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			domainName := fmt.Sprintf("%s.com", uuid.NewString())
			value := "1.2.3.4"
			defer func() {
				_, err := tt.domainHandler.Delete(domain.DeleteDomainRequest{Name: domainName})
				require.NoError(t, err)
			}()
			_, err := tt.domainHandler.Add(domain.AddDomainRequest{Name: domainName})
			require.NoError(t, err)

			_, err = tt.dnsHandler.Create(dns.CreateDnsRequest{
				Name:   uuid.NewString(),
				Value:  value,
				Domain: domainName,
				Type:   "A",
			})
			require.NoError(t, err)

			foundRes, err := tt.dnsHandler.List(dns.ListDnsRequest{Domain: domainName})
			require.NoError(t, err)
			require.Equal(t, 1, len(foundRes.Records))
			require.Equal(t, value, foundRes.Records[0].Value)
		})
	}
}

func TestListRecords(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			domainName := fmt.Sprintf("%s.com", uuid.NewString())
			value := "1.2.3.4"
			defer func() {
				_, err := tt.domainHandler.Delete(domain.DeleteDomainRequest{Name: domainName})
				require.NoError(t, err)
			}()
			_, err := tt.domainHandler.Add(domain.AddDomainRequest{Name: domainName})
			require.NoError(t, err)

			_, err = tt.dnsHandler.Create(dns.CreateDnsRequest{
				Name:   uuid.NewString(),
				Value:  value,
				Domain: domainName,
				Type:   "A",
			})
			require.NoError(t, err)

			foundRes, err := tt.dnsHandler.List(dns.ListDnsRequest{Domain: domainName})
			require.NoError(t, err)
			require.Equal(t, 1, len(foundRes.Records))
			require.Equal(t, value, foundRes.Records[0].Value)
		})
	}

}

func TestDeleteDomain(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			domainName := fmt.Sprintf("%s.com", uuid.NewString())
			value := "1.1.1.1"
			defer func() {
				_, err := tt.domainHandler.Delete(domain.DeleteDomainRequest{Name: domainName})
				require.NoError(t, err)
			}()
			_, err := tt.domainHandler.Add(domain.AddDomainRequest{Name: domainName})
			require.NoError(t, err)

			res, err := tt.dnsHandler.Create(dns.CreateDnsRequest{
				Name:   uuid.NewString(),
				Value:  value,
				Domain: domainName,
				Type:   "A",
			})
			require.NoError(t, err)
			_, err = tt.dnsHandler.Delete(dns.DeleteDnsRequest{Domain: domainName, Record: res.Uid})
			require.NoError(t, err)
			foundRes, err := tt.dnsHandler.List(dns.ListDnsRequest{Domain: domainName})
			require.NoError(t, err)
			require.Equal(t, 0, len(foundRes.Records))
		})
	}

}
