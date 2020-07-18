package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

// rsa公钥加密
func PublicRSAEncrypt(str, key string) (string, error) {
	// pem解码
	block, _ := pem.Decode([]byte(key))
	// x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(str))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// rsa私钥解密
func PrivateRSADecrypt(str, key string) (string, error) {
	cipherText, _ := base64.StdEncoding.DecodeString(str)
	// pem解码
	block, _ := pem.Decode([]byte(key))
	// x509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		return "", err
	}
	return string(plainText[:]), nil

}
