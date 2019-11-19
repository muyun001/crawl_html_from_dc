package get_response_html

import (
	"crypto/md5"
	"fmt"
)

// UniqueMd5 直接计算下载中心那边的unique_md5
// 默认逻辑unique_md5是从接口返回的，这个是原先python下载中心库中的生成逻辑
func UniqueMd5(url, uniqueKey string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(url+uniqueKey)))
}
