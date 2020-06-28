package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 加密  aes-128-cbc
func Encrypt(key string, iv string, str string) (string, error) {
	encodeKey := []byte(key)
	encodeIv := []byte(iv)
	encodeStr := []byte(str)

	//获取block块
	block, err := aes.NewCipher(encodeKey)
	if err != nil {
		return "", err
	}
	//补码
	encodeStr = PKCS7Padding(encodeStr, block.BlockSize())
	//加密模式
	blockMode := cipher.NewCBCEncrypter(block, encodeIv)
	//创建明文长度的数组
	dCrypts := make([]byte, len(encodeStr))
	//加密明文
	blockMode.CryptBlocks(dCrypts, encodeStr)
	return base64.StdEncoding.EncodeToString(dCrypts), nil
}

//补码
func PKCS7Padding(origData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(origData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padtext...)
}

func Decrypt(key string, iv string, str string) (string, error) {
	decodeKey := []byte(key)
	decodeIv := []byte(iv)
	strByte, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	//获取block块
	block, err := aes.NewCipher(decodeKey)
	if err != nil {
		return "", err
	}
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, decodeIv)
	//创建明文长度的数组
	decrypted := make([]byte, len(strByte))
	//加密明文
	blockMode.CryptBlocks(decrypted, strByte)
	//去掉字符
	decrypted = PKCS7upPadding(decrypted)
	return string(decrypted), nil
}

// 去掉字符
func PKCS7upPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
