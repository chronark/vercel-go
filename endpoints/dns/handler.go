package dns

import (
	"fmt"

	"github.com/chronark/vercel-go/api"
)

type DnsHandler struct {
	vercelClient api.VercelClient
	teamid       string
}

func New(vercelClient api.VercelClient, teamid string) *DnsHandler {
	return &DnsHandler{
		vercelClient,
		teamid,
	}
}

func (h *DnsHandler) Create(req CreateDnsRequest) (res CreateDnsResponse, err error) {
	apiRequest := api.NewApiRequest("POST", fmt.Sprintf("/v2/domains/%s/records", req.Domain), &res)
	apiRequest.Body = req
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return CreateDnsResponse{}, fmt.Errorf("Unable to create dns entry: %w", err)
	}
	return res, nil
}
func (h *DnsHandler) List(req ListDnsRequest) (res ListDnsResponse, err error) {
	apiRequest := api.NewApiRequest("GET", fmt.Sprintf("/v4/domains/%s/records", req.Domain), &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	if req.Limit != 0 {
		apiRequest.Query.Add("limit", fmt.Sprintf("%d", req.Limit))
	}
	if req.Until != 0 {
		apiRequest.Query.Add("until", fmt.Sprintf("%d", req.Until))
	}
	if req.Since != 0 {
		apiRequest.Query.Add("since", fmt.Sprintf("%d", req.Since))
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return ListDnsResponse{}, fmt.Errorf("Unable to list dns entries: %w", err)
	}
	return res, nil

}

func (h *DnsHandler) Delete(req DeleteDnsRequest) (res DeleteDnsResponse, err error) {
	apiRequest := api.NewApiRequest("DELETE", fmt.Sprintf("/v2/domains/%s/records/%s", req.Domain, req.Record), &res)
	if h.teamid != "" {
		apiRequest.Query.Add("teamId", h.teamid)
	}
	err = h.vercelClient.Call(apiRequest)
	if err != nil {
		return DeleteDnsResponse{}, fmt.Errorf("Unable to delete dns entry: %w", err)
	}
	return res, nil
}
