package jobs

import (
	"crawl_html_from_dc/services/api"
	"crawl_html_from_dc/services/build_dc_request"
)

// 构建下载中心请求
func BuildDcRequest(request *api.SendRequest) (*build_dc_request.DcSetTaskRequest, error) {
	dcRequest, err := build_dc_request.BuildRequest(request)
	if err != nil {
		return dcRequest, err
	}
	return dcRequest, nil
}
