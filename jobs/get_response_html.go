package jobs

import (
	"crawl_html_from_dc/services/api"
	"crawl_html_from_dc/services/get_response_html"
)

func GetResponseHtml(request *api.SendRequest) (string, error) {
	dcResponseBytes, err := get_response_html.GetDcResult(request)
	if err != nil {
		return "", err
	}

	if dcResponseBytes == nil {
		return "", nil
	}

	html, err := get_response_html.GetHtmlFromDcResponse(dcResponseBytes)
	if err != nil {
		return "", err
	}

	return html, nil
}
