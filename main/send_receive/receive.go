package send_receive

import (
	"crawl_html_from_dc/jobs"
	"crawl_html_from_dc/services/api"
	"crawl_html_from_dc/services/get_response_html"
	"crawl_html_from_dc/services/send_dc_request"
	"crawl_html_from_dc/settings"
	"errors"
	"fmt"
	"time"
)

// 异步获取下载中心结果
func AsynReceive(request *api.SendRequest) (*get_response_html.DcApiResponse, error) {
	dcResponseBytes, err := get_response_html.GetDcResult(request)
	if err != nil {
		return &get_response_html.DcApiResponse{}, err
	}

	dcGetResultResponse, err := get_response_html.UnmarshalDcResponse(dcResponseBytes)
	if err != nil {
		return &get_response_html.DcApiResponse{}, err
	}

	return dcGetResultResponse, nil
}

// 同步获取下载中心结果
func SyncReceive(request *api.SendRequest) (string, error) {
	var html string
	var err error
	startTimeToGetHtml := time.Now()

	for {
		html, err = jobs.GetResponseHtml(request)
		if err != nil {
			if newIp, err := send_dc_request.SendResetIp(); err == nil {
				settings.DcApi = "http://" + newIp
				html, err = jobs.GetResponseHtml(request)
				if err != nil {
					return "", errors.New(fmt.Sprintf("获取返回结果出错，err: %s", err.Error()))
				}
			} else {
				return "", errors.New(fmt.Sprintf("获取查询结果时重置接口IP出错：%s", err.Error()))
			}
		}

		if html != "" {
			break
		}

		time.Sleep(time.Second * 2)
		if time.Now().After(startTimeToGetHtml.Add(time.Minute * 10)) {
			return "", errors.New("获取结果超时")
		}
	}

	return html, nil
}
