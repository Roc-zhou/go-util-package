package wx

import (
	"crypto/sha1"
	"fmt"
	"strconv"

	"github.com/Roc-zhou/go-util-package/str"
	"github.com/Roc-zhou/go-util-package/time"
)

func WxSign(jsapi_ticket string, url string) map[string]interface{} {
	nonceStr := str.CreateNonceStr(16)
	timestamp := strconv.Itoa(int(time.GetCurStamp()))
	// 顺序必须有序
	string1 := "jsapi_ticket=" + jsapi_ticket + "&noncestr=" + nonceStr + "&timestamp=" + timestamp + "&url=" + url
	ret := make(map[string]interface{})
	ret["jsapi_ticket"] = jsapi_ticket
	ret["nonceStr"] = nonceStr
	ret["timestamp"] = timestamp
	ret["url"] = url
	ret["signature"] = fmt.Sprintf("%x", sha1.Sum([]byte(string1)))
	return ret
}
