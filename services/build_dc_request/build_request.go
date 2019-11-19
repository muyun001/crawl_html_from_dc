package build_dc_request

import (
	"crawl_html_from_dc/settings"
	"encoding/json"
)

func BuildRequest(url, cookie, userAgent string) (*DcSetTaskRequest, error) {
	headersStr, err := getHeaders(cookie, userAgent)
	if err != nil {
		return &DcSetTaskRequest{}, err
	}

	cookieStr, err := getConfig()
	if err != nil {
		return &DcSetTaskRequest{}, err
	}

	urlsStr, err := getUrls(url)
	if err != nil {
		return &DcSetTaskRequest{}, err
	}

	dcRequest := &DcSetTaskRequest{
		UserID:  settings.DcUserId,
		Headers: headersStr,
		Config:  cookieStr,
		Urls:    urlsStr,
	}

	return dcRequest, nil
}

func getHeaders(cookie, userAgent string) (string, error) {
	header := &DcHeaders{
		Cookie:    cookie,
		UserAgent: userAgent,
	}

	headerByte, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	return string(headerByte), nil
}

func getConfig() (string, error) {
	config := &DcConfig{
		Redirect: DcRedirectFalse,
		Priority: DcPriorityMiddle,
	}

	configByte, err := json.Marshal(config)
	if err != nil {
		return "", err
	}

	return string(configByte), nil
}

func getUrls(url string) (string, error) {
	var urls = []*DcUrl{}
	urls = append(urls, &DcUrl{
		Url:       url,
		Type:      DcResponseTypeHtml,
		UniqueKey: DcUniqueKey(url),
	})

	urlsByte, err := json.Marshal(urls)
	if err != nil {
		return "", err
	}

	return string(urlsByte), nil
}
