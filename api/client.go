package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type VercelClient interface {
	Call(method string, path string, body interface{}) (*http.Response, error)
}

type vercelClient struct {
	url        string
	httpClient HTTPClient
	userAgent  string
	token      string
}

type NewClientConfig struct {
	// Requests to the Vercel API must provide an API token.
	// You can get one here: https://vercel.com/account/tokens
	Token string

	// Optionally set a userAgent that is sent with every request.
	// Defaults to `chronark/vercel-go`
	UserAgent string

	// Use a custom HTTPClient.
	// This is mainly used for testing
	// Omit to use the default `net/http` implementation
	HTTPClient HTTPClient
}

func New(config NewClientConfig) VercelClient {
	if config.UserAgent == "" {
		config.UserAgent = "chronark/vercel-go"
	}
	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{}
	}
	return &vercelClient{
		url:        "https://api.vercel.com",
		httpClient: config.HTTPClient,
		userAgent:  config.UserAgent,
		token:      config.Token,
	}
}

// JSON marshal the body if present
func marshalBody(body interface{}) (io.Reader, error) {
	var payload io.Reader = nil
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		payload = bytes.NewBuffer(b)
	}
	return payload, nil
}

// Perform a request and return its response
func (c *vercelClient) Call(method string, path string, body interface{}) (*http.Response, error) {
	payload, err := marshalBody(body)
	if err != nil {
		return nil, fmt.Errorf("Unable to marshal request body: %w", err)
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.url, path), payload)
	if err != nil {
		return nil, fmt.Errorf("Unable to create request: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable to perform request: %w", err)
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var responseBody map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&responseBody)
		if err != nil {
			return nil, fmt.Errorf("Unable to decode response body of bad response: %s: %w", res.Status, err)
		}

		// Try to prettyprint the response body
		// If that is not possible we return the raw body
		pretty, err := json.MarshalIndent(responseBody, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("Response returned status code %d: %+v", res.StatusCode, responseBody)
		}
		return nil, fmt.Errorf("Response returned status code %d: %+v", res.StatusCode, string(pretty))
	}

	return res, nil

}
