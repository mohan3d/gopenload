package openload

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
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
	return nil, c.get("/account/info", nil)
}

func (c *Client) getAPIURL(p string, q map[string]string) (string, error) {
	u, err := url.Parse(c.api)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(u.Path, p)

	params := url.Values{}
	params.Add("login", c.login)
	params.Add("key", c.key)
	if q != nil {
		for k, v := range q {
			params.Add(k, v)
		}
	}
	u.RawQuery = params.Encode()

	return u.String(), nil
}

func (c *Client) get(p string, q map[string]string) error {
	return nil
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
