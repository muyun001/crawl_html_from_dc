package get_response_html

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

func GetHtmlFromDcResponse(dcResponseBytes []byte) (string, error) {
	dcGetResultResponse, err := UnmarshalDcResponse(dcResponseBytes)
	if err != nil {
		return "", err
	}

	if dcGetResultResponse.Status != DcCrawlSucceed {
		return "", errors.New(fmt.Sprintf("下载中心抓取数据失败，错误信息：%s", dcGetResultResponse.Msg))
	}

	rDataMap, err := ResponseRDataMap(dcGetResultResponse.RData)
	if err != nil {
		return "", err
	}

	html, err := DecodeHtml(rDataMap)
	if err != nil {
		return "", err
	}

	return html, nil
}

func UnmarshalDcResponse(dcResponseBytes []byte) (*DcApiResponse, error) {
	dcGetResultResponse := &DcApiResponse{}
	err := json.Unmarshal(dcResponseBytes, &dcGetResultResponse)
	if err != nil {
		return dcGetResultResponse, err
	}
	return dcGetResultResponse, nil
}

func DecodeHtml(rDataMap map[string]RData) (string, error) {
	var html string
	for uniqueMd5 := range rDataMap {
		if rDataMap[uniqueMd5].Status == 3 {
			return "", errors.New("下载中心抓取失败")
		}

		resultBytes, err := base64.StdEncoding.DecodeString(rDataMap[uniqueMd5].Result)
		if err != nil {
			return "", err
		}
		html = string(resultBytes)
	}
	return html, nil
}
