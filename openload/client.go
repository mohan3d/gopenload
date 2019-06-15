package openload

import (
	"fmt"
	"net/http"
)

const (
	apiBaseURL = "https://api.openload.co"
	apiVersion = "1"
)

func buildAPIURL() string {
	return fmt.Sprintf("%s/%s", apiBaseURL, apiVersion)
}

// Client represents openload api client.
type Client struct {
	login      string
	key        string
	api        string
	httpClient *http.Client
}

// AccountInfo requests logged-in account info
// And returns AccountInfo object holding infos.
func (c *Client) AccountInfo() (*AccountInfo, error) {
	return nil, nil
}

// New creates new openload client an returns a reference.
func New(login, key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		login:      login,
		key:        key,
		api:        buildAPIURL(),
		httpClient: httpClient,
	}
}
