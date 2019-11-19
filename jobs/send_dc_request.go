package jobs

import (
	"crawl_html_from_dc/services/build_dc_request"
	"crawl_html_from_dc/services/send_dc_request"
	"crawl_html_from_dc/utils/json_utils"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"
)

// 发送请求到下载中心
func SendDcRequest(request *build_dc_request.DcSetTaskRequest) (string, error) {
	var sendDcRequestMap map[string]string
	_ = json_utils.StructToStringMap(request, &sendDcRequestMap)

	formData := url.Values{}
	for key, value := range sendDcRequestMap {
		formData.Set(key, value)
	}

	body, err := send_dc_request.SendDcRequest(formData)
	if err != nil {
		time.Sleep(time.Second * 2)
		body, err = send_dc_request.SendDcRequest(formData)
		if err != nil {
			return "", err
		}
	}

	dcSendRequestResponse := &send_dc_request.DcApiResponse{}
	_ = json.Unmarshal(body, &dcSendRequestResponse)

	if strings.Contains(string(body), "params error") {
		return "", errors.New("发送请求到下载中心时参数错误")
	}

	uniqueMd5 := send_dc_request.ResponseUniqueMd5(dcSendRequestResponse.RData)
	if strings.Contains(string(body), "task insert error") {
		return uniqueMd5, errors.New("重复插入数据")
	}

	return uniqueMd5, nil
}
