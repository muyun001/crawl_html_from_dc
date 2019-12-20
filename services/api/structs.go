package api

type SendRequest struct {
	Url     string  `json:"url"`
	Headers Headers `json:"headers"`
}

type Headers struct {
	Accept                  string `json:"accept"`
	AcceptEncoding          string `json:"accept_encoding"`
	AcceptLanguage          string `json:"accept_language"`
	CacheControl            string `json:"cache_control"`
	Connection              string `json:"connection"`
	Cookie                  string `json:"cookie"`
	Host                    string `json:"host"`
	SecFetchMode            string `json:"sec_fetch_mode"`
	SecFetchSite            string `json:"sec_fetch_site"`
	SecFetchUser            string `json:"sec_fetch_user"`
	UpgradeInsecureRequests string `json:"upgrade_insecure_requests"`
	UserAgent               string `json:"user_agent"`
}
