package api

type SendRequest struct {
	Url     string  `json:"url"`
	Headers Headers `json:"headers"`
}

type Headers struct {
	UserAgent string `json:"user_agent"`
	Cookie    string `json:"cookie"`
}
