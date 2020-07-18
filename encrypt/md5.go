package encrypt

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	m := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", m)
}
