package build_dc_request

type DcHeaders struct {
	UserAgent string `json:"user_agent"`
	Cookie    string `json:"cookie"`
}

type DcConfig struct {
	Redirect int `json:"redirect"`
	Priority int `json:"priority"`
}

type DcUrl struct {
	Url       string `json:"url"`
	Type      int    `json:"type"`
	UniqueKey string `json:"unique_key"`
}

type DcSetTaskRequest struct {
	UserID  string `json:"user_id"`
	Headers string `json:"headers"`
	Config  string `json:"config"`
	Urls    string `json:"urls"`
}
