package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// 生成RSA私钥和公钥
func GetRSAkey(bits int) error {
	// 私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	// 通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//用pem格式编码
	//创建文件保存私钥
	privateFile, err := os.Create("privateKey.pem")
	if err != nil {
		return err
	}
	defer privateFile.Close()
	privateBlock := pem.Block{
		Type:  "Roc rsa private key",
		Bytes: x509PrivateKey,
	}
	// 写入文件
	if err = pem.Encode(privateFile, &privateBlock); err != nil {
		return err
	}

	// 获取公钥，保存文件
	publicKey := privateKey.PublicKey
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	publicFile, err := os.Create("publicKey.pem")
	if err != nil {
		return err
	}
	defer publicFile.Close()
	piblicBlock := pem.Block{
		Type:  "Roc rsa public key",
		Bytes: x509PublicKey,
	}
	// 写入文件
	if err = pem.Encode(publicFile, &piblicBlock); err != nil {
		return err
	}
	return nil
}
