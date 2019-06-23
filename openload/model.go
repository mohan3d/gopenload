package openload

// AccountInfoResponse represents account info response.
type AccountInfoResponse struct {
	Extid       string      `json:"extid"`
	Email       string      `json:"email"`
	SignupAt    string      `json:"signup_at"`
	StorageLeft int         `json:"storage_left"`
	StorageUsed interface{} `json:"storage_used"`
	Traffic     struct {
		Left    int `json:"left"`
		Used24H int `json:"used_24h"`
	} `json:"traffic"`
	Balance interface{} `json:"balance"`
}

// DownloadTicketResponse represents download ticket response.
type DownloadTicketResponse struct {
	Ticket     string      `json:"ticket"`
	CaptchaURL interface{} `json:"captcha_url"`
	CaptchaW   interface{} `json:"captcha_w"`
	CaptchaH   interface{} `json:"captcha_h"`
	WaitTime   int         `json:"wait_time"`
	ValidUntil string      `json:"valid_until"`
}

// DownloadLinkResponse represents download link response.
type DownloadLinkResponse struct {
	Name        string      `json:"name"`
	Size        interface{} `json:"size"`
	Sha1        string      `json:"sha1"`
	ContentType string      `json:"content_type"`
	UploadAt    string      `json:"upload_at"`
	URL         string      `json:"url"`
	Token       string      `json:"token"`
}

// FileInfoResponse represents single file info response.
type FileInfoResponse struct {
	ID          string      `json:"id"`
	Status      int         `json:"status"`
	Name        interface{} `json:"name"`
	Size        interface{} `json:"size"`
	Sha1        interface{} `json:"sha1"`
	ContentType interface{} `json:"content_type"`
}

// FilesInfoResponse represents multiple files info response.
type FilesInfoResponse map[string]FileInfoResponse

// UploadURLResponse represents upload url response.
type UploadURLResponse struct {
	URL        string `json:"url"`
	ValidUntil string `json:"valid_until"`
}

// UploadResponse represents upload response.
type UploadResponse struct {
	ContentType string `json:"content_type"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Sha1        string `json:"sha1"`
	Size        string `json:"size"`
	URL         string `json:"url"`
}

// RemoteUploadResponse represents remote upload response.
type RemoteUploadResponse struct {
	ID       string `json:"id"`
	Folderid string `json:"folderid"`
}

// RemoteUploadStatusResponse represents single remote upload status.
type RemoteUploadStatusResponse struct {
	ID          interface{} `json:"id"`
	Remoteurl   string      `json:"remoteurl"`
	Status      string      `json:"status"`
	BytesLoaded interface{} `json:"bytes_loaded"`
	BytesTotal  interface{} `json:"bytes_total"`
	Folderid    string      `json:"folderid"`
	Added       string      `json:"added"`
	LastUpdate  string      `json:"last_update"`
	Extid       interface{} `json:"extid"`
	URL         interface{} `json:"url"`
}

// RemoteUploadsStatusResponse represents all remote uploads status.
type RemoteUploadsStatusResponse map[string]RemoteUploadStatusResponse

// ListFolderResponse represents list folder response.
type ListFolderResponse struct {
	Folders []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"folders"`
	Files []struct {
		Name          string `json:"name"`
		Sha1          string `json:"sha1"`
		Folderid      string `json:"folderid"`
		UploadAt      string `json:"upload_at"`
		Status        string `json:"status"`
		Size          string `json:"size"`
		ContentType   string `json:"content_type"`
		DownloadCount string `json:"download_count"`
		Cstatus       string `json:"cstatus"`
		Link          string `json:"link"`
		Linkextid     string `json:"linkextid"`
	}
}

// RenameFolderResponse represents rename folder response either true or false.
type RenameFolderResponse bool

// RenameFileResponse represents rename file response either true or false.
type RenameFileResponse bool

// DeleteFileResponse represents delete file response either true or false.
type DeleteFileResponse bool

// ConvertFileResponse represents conver file response either true or false.
type ConvertFileResponse bool

// RunningConversionsResponse represents pending conversions response.
type RunningConversionsResponse []struct {
	Name       string  `json:"name"`
	ID         string  `json:"id"`
	Status     string  `json:"status"`
	LastUpdate string  `json:"last_update"`
	Progress   float64 `json:"progress"`
	Retries    string  `json:"retries"`
	Link       string  `json:"link"`
	Linkextid  string  `json:"linkextid"`
}

// SplashImageResponse represents splash image response.
type SplashImageResponse string
