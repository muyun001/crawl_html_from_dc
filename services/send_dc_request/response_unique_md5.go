package send_dc_request

import "regexp"

// ResponseUniqueMd5 从Response的rdata中提取unique_md5
func ResponseUniqueMd5(rData string) string {
	var uniqueMd5 string
	re := regexp.MustCompile(`\["(.*?)"\]`)
	subMatch := re.FindStringSubmatch(rData)
	if len(subMatch) == 2 {
		uniqueMd5 = subMatch[1]
	}

	return uniqueMd5
}
