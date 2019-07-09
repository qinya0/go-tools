package crypto

import "bytes"

// pkcs7 补码
func pKCS7Padding(data []byte, blockSize int) []byte {
	paddingLength := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)
	return append(data, padText...)
}
// pkcs7 去码
func pKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unPaddingLength := int(data[length-1])
	return data[:(length - unPaddingLength)]
}
