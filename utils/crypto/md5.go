package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var (
	DataNilError = fmt.Errorf("data is nil err")
)

type Md5 struct {}

func (m *Md5) Encrypt(data string) (reStr string, err error) {
	if data == "" {
		err = DataNilError
		return
	}
	h := md5.New()
	h.Write([]byte(data))
	reStr = hex.EncodeToString(h.Sum(nil))

	return
}

func (m *Md5) Decrypt(data string) (reStr string, err error) {
	var (
		reByte []byte
	)
	if data == "" {
		err = DataNilError
		return
	}
	h := md5.New()
	h.Write([]byte(data))
	reByte, err = hex.DecodeString(data)
	if err != nil {
		return
	}
	reStr = string(reByte)

	return
}

func md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return  hex.EncodeToString(h.Sum(nil))
}