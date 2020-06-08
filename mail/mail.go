package mail

import (
	"strings"

	"gopkg.in/gomail.v2"
)

type Params struct {
	Host        string
	Port        int
	Username    string // 发件人 & 用户名
	Password    string
	GetMailUser string // 收件人 多个 ｜
	MailTheme   string // 邮件主题
	Body        string // 邮件正文
}

func Send(opt *Params) {
	m := gomail.NewMessage()
	// 设置发件人
	m.SetHeader("From", opt.Username)
	// 收件人
	mailTo := strings.Split(opt.GetMailUser, "|")
	m.SetHeader("To", mailTo...)
	//设置邮件主题
	m.SetHeader("Subject", opt.MailTheme)
	// 邮件正文
	m.SetBody("text/html", opt.Body)

	d := gomail.NewDialer(opt.Host, opt.Port, opt.Username, opt.Password)

	// 发送
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
