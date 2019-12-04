package get_response_html

import (
	"crawl_html_from_dc/services/build_dc_request"
	"crawl_html_from_dc/settings"
	"crawl_html_from_dc/utils/json_utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/panwenbin/ghttpclient"
	"github.com/panwenbin/ghttpclient/header"
	"net/url"
	"strings"
	"time"
)

const GET_DC_GET_RESULT = "/download/getResult"

// 获取查询结果
func GetDcResult(requestUrl string) ([]byte, error) {
	configStr, err := getConfig()
	if err != nil {
		return nil, err
	}

	urlsStr, err := getUrls(requestUrl)
	if err != nil {
		return nil, err
	}

	dcGetResultRequest := &DcGetResultRequest{
		UserID: settings.DcUserId,
		Config: configStr,
		Urls:   urlsStr,
	}

	formData := url.Values{}
	err = json_utils.StructToFormData(dcGetResultRequest, &formData)
	if err != nil {
		return nil, err
	}

	body, err := getResult(formData)
	if err != nil {
		time.Sleep(time.Second * 2)
		body, err = getResult(formData)
		if err != nil {
			return nil, err
		}
	}

	//if !strings.Contains(string(body), "result") {
	//	return body, nil
	//}

	return body, nil
}

func getConfig() (string, error) {
	dcConfig := build_dc_request.DcConfig{
		Redirect: build_dc_request.DcRedirectFalse,
		Priority: build_dc_request.DcPriorityMiddle,
	}
	configJsonBytes, err := json.Marshal(dcConfig)
	if err != nil {
		return "", err
	}

	return string(configJsonBytes), nil
}

func getUrls(url string) (string, error) {
	dcUrls := make([]DcUrl, 0)
	uniqueKey := build_dc_request.DcUniqueKey(url)
	dcUrls = append(dcUrls, DcUrl{
		Url:       url,
		Type:      build_dc_request.DcResponseTypeHtml,
		UniqueKey: uniqueKey,
		UniqueMd5: UniqueMd5(url, uniqueKey),
	})

	urlsJsonBytes, err := json.Marshal(dcUrls)
	if err != nil {
		return "", err
	}

	return string(urlsJsonBytes), nil
}

func getResult(formData url.Values) ([]byte, error) {
	apiUrl := settings.DcApi + GET_DC_GET_RESULT
	client := ghttpclient.NewClient().Timeout(time.Second * 30).Url(apiUrl).
		Headers(nil).Body(strings.NewReader(formData.Encode())).
		ContentType(header.CONTENT_TYPE_FORM_URLENCODED).Post()

	res, err := client.Response()
	if err != nil {
		return nil, err
	}

	body, err := client.ReadBodyClose()
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 200 {
		return body, errors.New(fmt.Sprintf("获取下载中心结果时状态码错误: %d", res.StatusCode))
	}

	return body, nil
}
