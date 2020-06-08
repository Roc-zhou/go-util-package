package go_util_package

import (
	"testing"

	"github.com/Roc-zhou/go-util-package/mail"
)

func TestMail(t *testing.T) {
	// 发送邮件
	var params = &mail.Params{
		Host:        "smtp.exmail.qq.com",
		Port:        465,
		Username:    "",
		Password:    "",
		GetMailUser: "1137938565@qq.com", // 收件人
		MailTheme:   "服务错误",              // 邮件主题
		Body:        "hello mail",        // 邮件正文
	}
	mail.Send(params)
}
