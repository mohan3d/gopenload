package openload

import (
	"encoding/json"
	"errors"
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

func checkStatus(status int, msg string) error {
	if status != http.StatusOK {
		return errors.New(msg)
	}
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
// https://openload.co/api#accountinfos
func (c *Client) AccountInfo() (*AccountInfoResponse, error) {
	return nil, nil
}

// DownloadTicket requests download ticket for a specific file
// This ticket will be used for actual download.
// https://openload.co/api#download-ticket
func (c *Client) DownloadTicket(fileID string) (*DownloadTicketResponse, error) {
	return nil, nil
}

// DownloadLink requests direct download link for a specific file
// This step must be executed after getting ticket and captcha response
// From previous step (DownloadTicket).
// https://openload.co/api#download-getlink
func (c *Client) DownloadLink(fileID string, ticket string, captchaResponse string) (*DownloadLinkResponse, error) {
	return nil, nil
}

// FileInfo requests a single file info.
// https://openload.co/api#download-info
func (c *Client) FileInfo(fileID string) (*FileInfoResponse, error) {
	return nil, nil
}

// FilesInfo requests info for a list of files.
// https://openload.co/api#download-info
func (c *Client) FilesInfo(filesID []string) (FilesInfoResponse, error) {
	return nil, nil
}

// UploadLink requests an upload URL
// This URL will be used to perform actual upload.
// https://openload.co/api#upload
func (c *Client) UploadLink(folderID string, sha1 string, httponly bool) (*UploadURLResponse, error) {
	return nil, nil
}

// TODO: Upload.

// RemoteUpload adds a remote upload.
// https://openload.co/api#remoteul-add
func (c *Client) RemoteUpload(url string, folderID string) (*RemoteUploadResponse, error) {
	return nil, nil
}

// RemoteUploadStatus checks status of remote upload requests.
// https://openload.co/api#remoteul-check
func (c *Client) RemoteUploadStatus(limit int64, uploadID string) (RemoteUploadsStatusResponse, error) {
	return nil, nil
}

// ListFolder lists a folder content.
// https://openload.co/api#file-listfolder
func (c *Client) ListFolder(folderID string) (*ListFolderResponse, error) {
	return nil, nil
}

// RenameFolder renames existing folder.
// https://openload.co/api#file-renamefolder
func (c *Client) RenameFolder(folderID string, name string) (RenameFolderResponse, error) {
	return false, nil
}

// RenameFile renames existing file.
// https://openload.co/api#file-rename
func (c *Client) RenameFile(fileID string, name string) (RenameFolderResponse, error) {
	return false, nil
}

// DeleteFile deletes existing file.
// https://openload.co/api#file-delete
func (c *Client) DeleteFile(fileID string) (DeleteFileResponse, error) {
	return false, nil
}

// ConvertFile asks openload to convert media file.
// Usually conversions settings should be set in the UI
// https://openload.co/account#conversionsettings
// https://openload.co/api#convertingfiles
func (c *Client) ConvertFile(fileID string) (ConvertFileResponse, error) {
	return false, nil
}

// RunningConversions checks pending conversions status.
// https://openload.co/api#file-runningconverts
func (c *Client) RunningConversions(folderID string) (RunningConversionsResponse, error) {
	return nil, nil
}

// SplashImage request splash image direct url
// Usually it should be used with media fileID (movie, ...)
// https://openload.co/api#file-splash
func (c *Client) SplashImage(fileID string) (SplashImageResponse, error) {
	return "", nil
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
	var msg string

	err := json.NewDecoder(response).Decode(&data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(*data["status"], &status)
	if err != nil {
		return err
	}
	err = json.Unmarshal(*data["msg"], &msg)
	if err != nil {
		return err
	}
	err = checkStatus(status, msg)
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
