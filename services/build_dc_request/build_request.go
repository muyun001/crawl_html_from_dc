package build_dc_request

import (
	"crawl_html_from_dc/services/api"
	"crawl_html_from_dc/settings"
	"encoding/json"
)

func BuildRequest(request *api.SendRequest) (*DcSetTaskRequest, error) {
	headersStr, err := getHeaders(&request.Headers)
	if err != nil {
		return &DcSetTaskRequest{}, err
	}

	configStr, err := getConfig(&request.Config)
	if err != nil {
		return &DcSetTaskRequest{}, err
	}

	urlsStr, err := getUrls(request.Url.Url)
	if err != nil {
		return &DcSetTaskRequest{}, err
	}

	dcRequest := &DcSetTaskRequest{
		UserID:  settings.DcUserId,
		Headers: headersStr,
		Config:  configStr,
		Urls:    urlsStr,
	}

	return dcRequest, nil
}

func getHeaders(header *api.DcHeaders) (string, error) {
	headerByte, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	return string(headerByte), nil
}

func getConfig(config *api.DcConfig) (string, error) {
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
