package openload

import (
	"encoding/json"
	"fmt"
	"io"
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

func checkStatus(status int) error {
	return nil
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
	var info AccountInfo
	if err := c.get("/account/info", nil, &info); err != nil {
		return nil, err
	}
	return &info, nil
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

func processResponse(response io.Reader, result interface{}) error {
	var data map[string]*json.RawMessage
	var status int

	err := json.NewDecoder(response).Decode(&data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(*data["status"], &status)
	if err != nil {
		return err
	}
	err = checkStatus(status)
	if err != nil {
		return err
	}
	return json.Unmarshal(*data["result"], &result)
}

func (c *Client) get(p string, q map[string]string, result interface{}) error {
	u, err := c.getAPIURL(p, q)
	if err != nil {
		return err
	}
	response, err := http.Get(u)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return processResponse(response.Body, &result)
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
