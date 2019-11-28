package utils

import (
	"math/rand"
	"time"
)

func BaiduCookie() string {
	numStr := "0123456789"
	charStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numBytes := []byte(numStr)
	charBytes := []byte(charStr)

	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 21; i++ {
		result = append(result, numBytes[r.Intn(len(numBytes))])
	}
	for i := 0; i < 12; i++ {
		result = append(result, charBytes[r.Intn(len(charBytes))])
	}
	for i := len(result) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		result[i], result[num] = result[num], result[i]
	}

	return string(result[0:32])
}
