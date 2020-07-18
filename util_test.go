package main

import (
	"fmt"
	"testing"

	"github.com/Roc-zhou/go-util-package/encrypt"
	"github.com/Roc-zhou/go-util-package/mail"
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
	fmt.Println(encrypt.Md5(str))
}

func TestGetRSA(t *testing.T) {
	encrypt.GetRSAkey(2048)

	const (
		text       = "hello world"
		publickKey = `-----BEGIN rsa public key-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwWYMTKPGRFJDLQYr0g8i
OO5tPWnEbaDDLXEVa4kqhbRE6xAlflu5StDhMtPNdlMjGGIOFiQXfqHFGMF8CFMi
IO1VJ7VlXfDxFpa2bn9E8CJNrhfP9h60Ekw5E8g0FDFI3u+ckVf5uhXz8acu8ch+
pIUtCFD6QJ82bZyA/3WCkv6enwG58dFnVgtP2ZWrIuk4lrhS+rZEzbCyzv+kzdpB
mr7lgXGIKK0D2E3EM5GSYGxjhILCn8O9NppSuJT+mLpdrwBuPUTAcZFy5h0YqEcw
X87S8W+Hrxy1MrYQKTzDNZ0fGpqOQ7KmGMp2n2962UFAw9JGILFA7U7IeE6U3TVp
2wIDAQAB
-----END rsa public key-----`
		privateKey = `-----BEGIN rsa private key-----
MIIEpAIBAAKCAQEAwWYMTKPGRFJDLQYr0g8iOO5tPWnEbaDDLXEVa4kqhbRE6xAl
flu5StDhMtPNdlMjGGIOFiQXfqHFGMF8CFMiIO1VJ7VlXfDxFpa2bn9E8CJNrhfP
9h60Ekw5E8g0FDFI3u+ckVf5uhXz8acu8ch+pIUtCFD6QJ82bZyA/3WCkv6enwG5
8dFnVgtP2ZWrIuk4lrhS+rZEzbCyzv+kzdpBmr7lgXGIKK0D2E3EM5GSYGxjhILC
n8O9NppSuJT+mLpdrwBuPUTAcZFy5h0YqEcwX87S8W+Hrxy1MrYQKTzDNZ0fGpqO
Q7KmGMp2n2962UFAw9JGILFA7U7IeE6U3TVp2wIDAQABAoIBAQCD1Q5ZR+KlO/Yy
wNwqKsrHSDALBwgxOr8RQN67GRt1XPcFzVHhmqDqUQzR2vNZzz5DJsQ3b47ccWWr
hGkKO7EuBNphYROiP9X38fCVzgeuMZQGMpE+UpupRXA6/eQSR65G2cs+gFvo74IE
nlQv2N4LRAT4gEq2tlEh0udUPEQv8lVeBPh/34IU002T/66kJuLQUsVqZwdYpYSQ
fY9oP0gTGHpl/w0RDFef00ub8s2lkK7NDyejDuJn04zIjOEIvy58CKVy30bDHoOy
HxxhBb5Az9olr5R2ujwkskv0OWtkYUxKxStAOsgwliYsdfvWfuYjaCEhHhVHWyru
EZ7GCL6BAoGBANLPAL6WHgRSdcHwIyNc76GYq1HdMMJBl9bg5AM/pcDZl9odmuXs
ewypBJk/hblNvkAnoXshBP6Cs1ZJvA2kE9+eqRbirntlHbLdF22n4WdD5exSavfA
SOXVDNnwHVyWTnpYnuXpwYa1gk8I8Qz10EAqOVUmaDB3DEYPpgnJ1Zp7AoGBAOrb
majkk7t3ceuelAPiIwm7JR6QF7eFXVp4A9yo816NpP18PdrogwM0xTA6x65vyTHr
ZTZKjLLcEpo2nHcGu/lxow3kkihKjQoaT9blSsbQTV9LbJ0w+aMcQGwEZSReV+lZ
q306T1P5Ppp6m9vgAqpg7BYdcLGJCp/4ehOKnYAhAoGBALVSW7+Se4sYKjWACZk3
LN5/5Ivrhy7vIF0w3q50pmt0PYrcgAlYGJbRIiV2X9z9I7Em5Vx1EEihvwNvNJM/
F+D2JVaL76wvUxYv2SD2j4g0/KUCVwN2nr0hPYaPY5HpruLJKxNytoTZxpWYiU2u
eUtXhizZdjCJlUMF3rj4pAPVAoGAS7/8tb7T73k9IzQIpaAAkRjthggPvj+jtpRT
Go7bwDmLZ7707HBmIViZ6U5sLVUc3Z8BDBvLeb0Fuvu2R6XCZ1hBsS9x/NDe4M1P
xw45qhpxejUHmUO8oOFx3eUlAi/zxu4HI+L0xy7zBDxbnPWJpo6QVymzuW9sH+Cy
7y6cUkECgYBnkMB66XPYoe2v97OpnR4K9GC6NUg53NtfD4Pj+iqvbsAATWq9BGFk
CljZ9CmhK2biLVbS+Zo70hHZDpC5HK8mhTrkwjafuvXAjMfocJWDcoOOuqXal78I
SEdoG2OMA8Tb4x+VeqnEOXAabpXlmudT4iC7Wf9yvAYhCUhTFK3/Rw==
-----END rsa private key-----`
	)

	// 公钥加密
	cipherText, _ := encrypt.PublicRSAEncrypt(text, publickKey)
	fmt.Println("加密后：" + cipherText)
	// 私钥解密
	plainText, _ := encrypt.PrivateRSADecrypt(cipherText, privateKey)
	fmt.Println("原文：" + plainText)
}
