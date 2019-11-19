package build_dc_request

import (
	"crypto/md5"
	"fmt"
)

func DcUniqueKey(requestUrl string) string {
	//dateStr := time.Now().Format("2006-01-02")
	//sourceStr := fmt.Sprintf("%s%s", requestUrl, dateStr)
	return fmt.Sprintf("%x", md5.Sum([]byte(requestUrl)))
}
