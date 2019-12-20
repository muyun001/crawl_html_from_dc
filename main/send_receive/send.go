package send_receive

import (
	"crawl_html_from_dc/jobs"
	"crawl_html_from_dc/services/api"
	"crawl_html_from_dc/services/send_dc_request"
	"crawl_html_from_dc/settings"
	"errors"
	"fmt"
)

// 发送请求
func Send(request *api.SendRequest) error {
	// 构建下载中心请求
	dcRequest, err := jobs.BuildDcRequest(request)
	if err != nil {
		return errors.New(fmt.Sprintf("构建请求出错，err: %s", err.Error()))
	}

	// 发送任务到下载中心
	_, err = jobs.SendDcRequest(dcRequest) // uniqueMd5
	if err != nil {
		fmt.Println(err)
		if err.Error() == "重复插入任务" {
			return err
		}

		if newIp, err := send_dc_request.SendResetIp(); err == nil {
			settings.DcApi = "http://" + newIp
			_, err = jobs.SendDcRequest(dcRequest)
			if err != nil {
				return errors.New(fmt.Sprintf("发送请求出错，err: %s", err.Error()))
			}
		} else {
			return errors.New(fmt.Sprintf("发送请求时重置接口IP出错：%s", err.Error()))
		}
	}

	return nil
}
