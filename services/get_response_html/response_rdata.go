package get_response_html

import (
	"encoding/json"
	"strings"
)

// ResponseRDataMap 从Response的rdata中提取RData结构的Map，key为unique_md5
func ResponseRDataMap(resRData string) (map[string]RData, error) {
	resRData = strings.Replace(resRData, `"status": "3"`, `"status": 3`, -1)
	rDatas := make(map[string]RData, 0)
	err := json.Unmarshal([]byte(resRData), &rDatas)
	if err != nil {
		return nil, err
	}

	return rDatas, nil
}
