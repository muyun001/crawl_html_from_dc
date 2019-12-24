package api

type SendRequest struct {
	Url     DcUrl     `json:"url"`
	Headers DcHeaders `json:"headers"`
	Config  DcConfig  `json:"config"`
}

type DcUrl struct {
	Url       string `json:"url"`
	Type      int    `json:"type"`
	UniqueKey string `json:"unique_key"`
}

type DcHeaders struct {
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

type DcConfig struct {
	StoreType      int    `json:"store_type"`
	Redirect       int    `json:"redirect"`
	Priority       int    `json:"priority"`
	PostData       string `json:"post_data"`
	ConfDistrictId int    `json:"conf_district_id"`
	Single         int    `json:"single"`
	ExpireTime     int    `json:"expire_time"`
}
