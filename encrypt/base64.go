package encrypt

import (
	"encoding/base64"
	"fmt"
)

func StringToBase64(str string) string {
	data := []byte(str)
	return base64.StdEncoding.EncodeToString(data)
}

func Base64ToString(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Printf("base64 decode failure, error=[%v]\n", err)
	}
	return string(data)
}
