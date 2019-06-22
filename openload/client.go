package openload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
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
	var info AccountInfoResponse
	if err := c.get("/account/info", nil, &info); err != nil {
		return nil, err
	}
	return &info, nil
}

// DownloadTicket requests download ticket for a specific file
// This ticket will be used for actual download.
// https://openload.co/api#download-ticket
func (c *Client) DownloadTicket(fileID string) (*DownloadTicketResponse, error) {
	var ticket DownloadTicketResponse
	if err := c.get("/file/dlticket", map[string]string{"file": fileID}, &ticket); err != nil {
		return nil, err
	}
	return &ticket, nil
}

// DownloadLink requests direct download link for a specific file
// This step must be executed after getting ticket and captcha response
// From previous step (DownloadTicket).
// https://openload.co/api#download-getlink
func (c *Client) DownloadLink(fileID string, ticket string, captchaResponse string) (*DownloadLinkResponse, error) {
	var link DownloadLinkResponse
	if err := c.get("/file/dl", map[string]string{"file": fileID, "ticket": ticket, "captcha_response": captchaResponse}, &link); err != nil {
		return nil, err
	}
	return &link, nil
}

// FileInfo requests a single file info.
// https://openload.co/api#download-info
func (c *Client) FileInfo(fileID string) (*FileInfoResponse, error) {
	infos, err := c.FilesInfo([]string{fileID})
	if err != nil {
		return nil, err
	}
	if l := len(infos); l != 1 {
		return nil, fmt.Errorf("expected only one file info got %d", l)
	}
	for _, v := range infos {
		return &v, nil
	}
	return nil, errors.New("this error is not supposed to happen")
}

// FilesInfo requests info for a list of files.
// https://openload.co/api#download-info
func (c *Client) FilesInfo(filesID []string) (FilesInfoResponse, error) {
	var info FilesInfoResponse
	if err := c.get("/file/info", map[string]string{"file": strings.Join(filesID, ",")}, &info); err != nil {
		return nil, err
	}
	return info, nil
}

// UploadLink requests an upload URL
// This URL will be used to perform actual upload.
// https://openload.co/api#upload
func (c *Client) UploadLink(folderID string, sha1 string, httponly bool) (*UploadURLResponse, error) {
	var link UploadURLResponse
	if err := c.get("/file/ul", map[string]string{"folder": folderID, "sha1": sha1, "httponly": strconv.FormatBool(httponly)}, &link); err != nil {
		return nil, err
	}
	return &link, nil
}

// Upload uploads contesnt of filepath.
// https://openload.co/api#upload
func (c *Client) Upload(filePath string, folderID string, sha1 string, httponly bool) (*UploadResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return c.UploadFrom(file, "", folderID, sha1, httponly)
}

// UploadFrom uploads contents of a reader.
// https://openload.co/api#upload
func (c *Client) UploadFrom(r io.Reader, name string, folderID string, sha1 string, httponly bool) (*UploadResponse, error) {
	return nil, nil
}

// RemoteUpload adds a remote upload.
// https://openload.co/api#remoteul-add
func (c *Client) RemoteUpload(url string, folderID string) (*RemoteUploadResponse, error) {
	var remote RemoteUploadResponse
	params := map[string]string{"url": url}
	if folderID != "" {
		params["folder"] = folderID
	}
	if err := c.get("/remotedl/add", params, &remote); err != nil {
		return nil, err
	}
	return &remote, nil
}

// RemoteUploadStatus checks status of remote upload requests.
// https://openload.co/api#remoteul-check
func (c *Client) RemoteUploadStatus(limit int64, uploadID string) (RemoteUploadsStatusResponse, error) {
	var status RemoteUploadsStatusResponse
	params := make(map[string]string)
	if limit == -1 {
		limit = 5
	}
	params["limit"] = strconv.FormatInt(limit, 10)
	if uploadID != "" {
		params["id"] = uploadID
	}
	if err := c.get("/remotedl/status", params, &status); err != nil {
		return nil, err
	}
	return status, nil
}

// ListFolder lists a folder content.
// https://openload.co/api#file-listfolder
func (c *Client) ListFolder(folderID string) (*ListFolderResponse, error) {
	var list ListFolderResponse
	params := make(map[string]string)
	if folderID != "" {
		params["folder"] = folderID
	}
	if err := c.get("/file/listfolder", params, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// RenameFolder renames existing folder.
// https://openload.co/api#file-renamefolder
func (c *Client) RenameFolder(folderID string, name string) (RenameFolderResponse, error) {
	var renamed RenameFolderResponse
	if err := c.get("/file/renamefolder", map[string]string{"folder": folderID, "name": name}, &renamed); err != nil {
		return renamed, err
	}
	return renamed, nil
}

// RenameFile renames existing file.
// https://openload.co/api#file-rename
func (c *Client) RenameFile(fileID string, name string) (RenameFileResponse, error) {
	var renamed RenameFileResponse
	if err := c.get("/file/rename", map[string]string{"file": fileID, "name": name}, &renamed); err != nil {
		return renamed, err
	}
	return renamed, nil
}

// DeleteFile deletes existing file.
// https://openload.co/api#file-delete
func (c *Client) DeleteFile(fileID string) (DeleteFileResponse, error) {
	var deleted DeleteFileResponse
	if err := c.get("/file/delete", map[string]string{"file": fileID}, &deleted); err != nil {
		return deleted, err
	}
	return deleted, nil
}

// ConvertFile asks openload to convert media file.
// Usually conversions settings should be set in the UI
// https://openload.co/account#conversionsettings
// https://openload.co/api#convertingfiles
func (c *Client) ConvertFile(fileID string) (ConvertFileResponse, error) {
	var converted ConvertFileResponse
	if err := c.get("/file/convert", map[string]string{"file": fileID}, &converted); err != nil {
		return converted, err
	}
	return converted, nil
}

// RunningConversions checks pending conversions status.
// https://openload.co/api#file-runningconverts
func (c *Client) RunningConversions(folderID string) (RunningConversionsResponse, error) {
	var conversions RunningConversionsResponse
	params := make(map[string]string)
	if folderID != "" {
		params["folder"] = folderID
	}
	if err := c.get("/file/runningconverts", params, &conversions); err != nil {
		return nil, err
	}
	return conversions, nil
}

// SplashImage request splash image direct url
// Usually it should be used with media fileID (movie, ...)
// https://openload.co/api#file-splash
func (c *Client) SplashImage(fileID string) (SplashImageResponse, error) {
	var image SplashImageResponse
	if err := c.get("/file/getsplash", map[string]string{"file": fileID}, &image); err != nil {
		return "", err
	}
	return image, nil
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
