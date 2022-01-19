package domain

import (
	"fmt"

	"github.com/chronark/vercel-go/api"
)

type DomainHandler struct {
	vercelClient api.VercelClient
	teamid       string
}

func New(vercelClient api.VercelClient, teamid string) *DomainHandler {
	return &DomainHandler{
		vercelClient,
		teamid,
	}
}

// Retrieves a list of domains registered for the authenticating user. By default it returns the
// last 20 domains if no limit is provided.
func (h *DomainHandler) List(req ListDomainsRequest) (res ListDomainsResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v5/domains", &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}

	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return ListDomainsResponse{}, fmt.Errorf("Unable to fetch domains from vercel: %w", err)
	}
	return res, nil
}

// Register a new domain name with Vercel for the authenticating user. The
// field serviceType selects whether the domains are going to use zeit.world
// DNS or an external nameserver. In the latter case a CNAME/ALIAS record(s)
// are expected to point towards cname.vercel-dns.com.
//
// If an external nameserver is used the user must verify the domain name by
// creating a TXT record for _now subdomain containing a verification token
// provided as a POST result. After the record has been created, the user may
// retry the same POST and the endpoint shall return verified: true, if the
// domain was verified successfully.
func (h *DomainHandler) Add(req AddDomainRequest) (res AddDomainResponse, err error) {
	apiRequest := api.NewApiRequest("POST", "/v4/domains", &res)
	apiRequest.Body = req
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}

	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return AddDomainResponse{}, fmt.Errorf("Unable to add domain to vercel: %w", err)
	}
	return res, nil
}

// Initiate a domain transfer request from an external Registrar to Vercel.
func (h *DomainHandler) Transfer(req TransferDomainRequest) (res TransferDomainResponse, err error) {
	apiRequest := api.NewApiRequest("POST", "/v4/domains", &res)
	apiRequest.Body = req
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return TransferDomainResponse{}, fmt.Errorf("Unable to transfer domain to vercel: %w", err)
	}
	return res, nil
}

// Verify a domain after adding either the intended nameservers or DNS TXT
// verification record to the domain.
func (h *DomainHandler) Verify(req VerifyDomainRequest) (res VerifyDomainResponse, err error) {

	apiRequest := api.NewApiRequest("POST", fmt.Sprintf("/v4/domains/%s/verify", req.Name), &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return VerifyDomainResponse{}, fmt.Errorf("Unable to verify domain: %w", err)
	}
	return res, nil
}

// Get information for a single domain in an account or team.
func (h *DomainHandler) Get(req GetDomainRequest) (res GetDomainResponse, err error) {
	apiRequest := api.NewApiRequest("GET", fmt.Sprintf("/v4/domains/%s", req.Name), &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}

	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return GetDomainResponse{}, fmt.Errorf("Unable to get domain from vercel: %w", err)
	}
	return res, nil
}

// Get Auth Code for a Single Domain
func (h *DomainHandler) GetAuthCode(req GetAuthCodeRequest) (res GetAuthCodeResponse, err error) {
	apiRequest := api.NewApiRequest("GET", fmt.Sprintf("/v6/domains/%s/auth-code", req.Name), &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return GetAuthCodeResponse{}, fmt.Errorf("Unable to get auth code: %w", err)
	}
	return res, nil
}

// Delete a domain by name
func (h *DomainHandler) Delete(req DeleteDomainRequest) (res DeleteDomainResponse, err error) {
	apiRequest := api.NewApiRequest("DELETE", fmt.Sprintf("/v4/domains/%s", req.Name), &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}

	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return DeleteDomainResponse{}, fmt.Errorf("Unable to remove domain from vercel: %w", err)
	}
	return res, nil
}

// Check if a domain name may be available to buy or not
func (h *DomainHandler) CheckAvailability(req CheckAvailabilityRequest) (res CheckAvailabilityResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v4/domains/status", &res)
	apiRequest.Query.Add("name", req.Name)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return CheckAvailabilityResponse{}, fmt.Errorf("Unable to check availability: %w", err)
	}
	return res, nil
}

// Check the price to purchase a domain and how long a single purchase period is
func (h *DomainHandler) CheckPrice(req CheckPriceRequest) (res CheckPriceResponse, err error) {
	apiRequest := api.NewApiRequest("GET", "/v4/domains/price", &res)
	apiRequest.Query.Add("name", req.Name)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return CheckPriceResponse{}, fmt.Errorf("Unable to check price: %w", err)
	}
	return res, nil
}

// Purchase the specified domain
func (h *DomainHandler) Purchase(req PurchaseRequest) (res PurchaseResponse, err error) {
	apiRequest := api.NewApiRequest("POST", "/v4/domains/buy", &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)

	if err != nil {
		return PurchaseResponse{}, fmt.Errorf("Unable to purchase domain from vercel: %w", err)
	}
	return res, nil
}
