package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var (
	NilKeyError = fmt.Errorf("key is nil err")
)

type AesCBC struct {
	key string
}

func NewAesCBC(k string) *AesCBC {
	return &AesCBC{
		key: k,
	}
}

// jia mi
func (ac *AesCBC) Encrypt(data string) (reStr string, err error) {
	if ac.key == "" {
		err = NilKeyError
		return
	}
	reStr = string(aesCBCEncrypt([]byte(data), []byte(md5Encode(ac.key))))
	return
}
func aesCBCEncrypt(data, key []byte) (reStr []byte) {
	// 分组秘钥
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	// 补码
	oData := pKCS7Padding(data, blockSize)
	// get block mode
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// add []byte to save data
	cryted := make([]byte, len(oData))
	// encode
	blockMode.CryptBlocks(cryted, oData)
	// encode again by base64
	return []byte(base64.StdEncoding.EncodeToString(cryted))
}

// jie mi
func (ac *AesCBC) Decrypt(data string) (reStr string, err error) {
	if ac.key == "" {
		err = NilKeyError
		return
	}
	reStr = string(aesCBCDecrypt([]byte(data), []byte(md5Encode(ac.key))))
	return
}
func aesCBCDecrypt(data, key []byte) []byte {
	// decode base64
	oData, _ := base64.StdEncoding.DecodeString(string(data))

	// get block
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])

	// add []byte
	reData := make([]byte, len(oData))
	// decode
	blockMode.CryptBlocks(reData, oData)
	// 去码
	reData = pKCS7UnPadding(reData)

	return reData
}


