package send_dc_request

import (
	"crawl_html_from_dc/settings"
	"errors"
	"fmt"
	"github.com/panwenbin/ghttpclient"
	"github.com/panwenbin/ghttpclient/header"
	"net/url"
	"strings"
	"time"
)

const PostSendRequest = "/download/setTask"

// 发送请求
func SendDcRequest(formData url.Values) ([]byte, error) {
	apiUrl := settings.DcApi + PostSendRequest
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
		return body, errors.New(fmt.Sprintf("访问下载中心API发送请求时状态码错误: %d", res.StatusCode))
	}

	return body, nil
}
