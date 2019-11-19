package get_response_html

type DcSetTaskRequest struct {
	UserID  string `json:"user_id"`
	Headers string `json:"headers"`
	Config  string `json:"config"`
	Urls    string `json:"urls"`
}

type DcGetResultRequest struct {
	UserID string `json:"user_id"`
	Config string `json:"config"`
	Urls   string `json:"urls"`
}

type DcConfig struct {
	Redirect int `json:"redirect"`
	Priority int `json:"priority"`
}

type DcUrl struct {
	Url       string `json:"url"`
	Type      int    `json:"type"`
	UniqueKey string `json:"unique_key"`
	UniqueMd5 string `json:"unique_md5,omitempty"`
}

type DcApiResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	RData  string `json:"rdata"`
}

type RData struct {
	Status      int    `json:"status"`
	Code        int    `json:"code"`
	InterPro    string `json:"inter_pro"`
	Header      string `json:"header"`
	Result      string `json:"result"`
	RedirectUrl string `json:"redirect_url"`
}
