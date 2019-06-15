package openload

// AccountInfo represents account info response.
type AccountInfo struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Result struct {
		Extid       string `json:"extid"`
		Email       string `json:"email"`
		SignupAt    string `json:"signup_at"`
		StorageLeft int    `json:"storage_left"`
		StorageUsed string `json:"storage_used"`
		Traffic     struct {
			Left    int `json:"left"`
			Used24H int `json:"used_24h"`
		} `json:"traffic"`
		Balance int `json:"balance"`
	} `json:"result"`
}
