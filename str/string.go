package str

import (
	"math/rand"
	"time"
)

// 随机字符串
func CreateNonceStr(num int) string {
	str := "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	bytes := []byte(str)
	result := []byte{}
	// 0-61
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result[:])
}
