package send_dc_request

import (
	"errors"
	"fmt"
	"github.com/panwenbin/ghttpclient"
	"github.com/panwenbin/ghttpclient/header"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const POST_RESET_IP = "/adslGetIp"

const INNER = 1
const OUTER = 2

// dcIpType 根据环境变量IsDcOuter决定接口IP是内网还是外网
func dcIpType() int {
	switch settings.DcIsInner {
	case true:
		return INNER
	case false:
		return OUTER
	}

	return INNER
}

// SendResetIp 重新获取IP地址(网络请求)
// 返回ip:port
func SendResetIp() (string, error) {
	apiUrl := settings.DcApi + POST_RESET_IP
	formData := url.Values{}
	formData.Set("type", strconv.Itoa(dcIpType()))

	client := ghttpclient.NewClient().Timeout(time.Second * 30).Url(apiUrl).
		Headers(nil).Body(strings.NewReader(formData.Encode())).
		ContentType(header.CONTENT_TYPE_FORM_URLENCODED).Post()

	res, err := client.Response()
	if err != nil {
		return "", err
	}

	body, _ := client.ReadBodyClose()

	if res.StatusCode > 200 {
		return "", errors.New(fmt.Sprintf("unexpect status: %d", res.StatusCode))
	}

	return string(body), nil
}

// ApiResetIp 重新获取IP地址
// 包含循环尝试
func ApiResetIp() {
	for {
		newIp, err := SendResetIp()
		if err == nil {
			settings.DcApi = "http://" + newIp
			fmt.Printf("设置DcAPI为: %s\n", settings.DcApi)
			break
		}

		fmt.Println("重新获取IP地址失败, 5秒后重试")
		time.Sleep(time.Second * 5)
	}
}
