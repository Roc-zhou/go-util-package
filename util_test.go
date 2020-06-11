package main

import (
	"fmt"
	"testing"

	"github.com/Roc-zhou/go-util-package/mail"
	"github.com/Roc-zhou/go-util-package/md5"
	"github.com/Roc-zhou/go-util-package/time"
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

func TestTime(t *testing.T) {
	stamp1 := fmt.Sprintf("当前时间： %s", time.GetNow())
	stamp2 := fmt.Sprintf("当前时间(秒)： %d", time.GetCurStamp())
	stamp3 := fmt.Sprintf("当前时间(毫秒)： %d", time.GetCurStamps())
	fmt.Println(stamp1)
	fmt.Println(stamp2)
	fmt.Println(stamp3)
}

func TestMd5(t *testing.T) {
	var str string = "123"
	fmt.Println(md5.Md5(str))
}
