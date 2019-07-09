package pwd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/qinya0/go-tools/utils/crypto"
	"github.com/qinya0/go-tools/utils/file"
)

// a manager for pwd
type Manager struct {
	// a total pwd for manager other pwd
	password string
	// file for save data
	file     string
}

func NewManager(pwd, file string) *Manager {
	return &Manager{
		password: pwd,
		file:     file,
	}
}

// save one pwd to file
func (m *Manager) SaveOne(name, passwd, msg string) (err error) {
	var (
		encodePwd  string
		marshalPwd []byte
		saveStr    string
		pwd        *PWD
	)
	// encrypt
	encodePwd, err = crypto.NewAesCBC(m.password).Encrypt(passwd)
	if err != nil {
		fmt.Printf("[manager] Encrypt passwd(%s) by key(%s) fail", passwd, m.password)
		return
	}
	//fmt.Printf("[test] NewAesCBC(%s)  Encrypt(%s) get(%s) \n", m.password, passwd, encodePwd)
	// get pwd
	pwd = &PWD {
		Key: name,
		Pwd: encodePwd,
		Msg: msg,
	}
	marshalPwd, err = json.Marshal(pwd)
	if err != nil {
		fmt.Printf("[manager] json.Marshal pwd(%s) err:%s", pwd.String(), err.Error())
		return
	}
	//fmt.Printf("[test] marshalPwd(%s) byte(%v) \n", string(marshalPwd), marshalPwd)
	// encrypt again
	saveStr, err = crypto.NewAesCBC(m.password).Encrypt(string(marshalPwd))
	if err != nil {
		fmt.Printf("[manager] Encrypt pwd(%s) by key(%s) fail", pwd.String(), m.password)
		return
	}
	//fmt.Printf("[test] save pwd(%s) encrypt(%s)\n", pwd.String(), saveStr)
	// save to file
	err = file.SaveString(m.file, saveStr)
	if err != nil {
		fmt.Printf("[manager] file.SaveString pwd(%s) err:%s", pwd.String(), err.Error())
		return
	}

	return
}

// get all pwd ferom file
func (m *Manager) GetAll() (result []PWD, err error) {
	var pwd PWD
	// open file
	strs, err := file.GetDataFromFile(m.file)
	if err != nil {
		fmt.Printf("[manager] GetDataFromFile(%s) err:%s", m.file, err.Error())
		return
	}
	for _, s := range strs {
		//fmt.Printf("[test] showall strs(%s) \n", s)
		if strings.TrimSpace(s) == "" {
			continue
		}
		pwd, err = m.getPWD(s)
		if err != nil {
			fmt.Printf("[manager] getPWD(%s) err:%s", s, err.Error())
			return
		}

		result = append(result, pwd)
	}
	// decode
	return
}

func (m *Manager) GetOne(key string) (result []PWD, err error) {
	var pwd PWD
	// open file
	strs, err := file.GetDataFromFile(m.file)
	if err != nil {
		fmt.Printf("[manager] GetDataFromFile(%s) err:%s", m.file, err.Error())
		return
	}
	for _, s := range strs {
		//fmt.Printf("[test] showall strs(%s) \n", s)
		if strings.TrimSpace(s) == "" {
			continue
		}
		pwd, err = m.getPWD(s)
		if err != nil {
			fmt.Printf("[manager] getPWD(%s) err:%s", s, err.Error())
			return
		}
		if pwd.Key != key {
			continue
		}
		result = append(result, pwd)
	}
	// decode
	return
}

func (m *Manager) getPWD(str string) (pwd PWD, err error) {
	var (
		DecryptStr string
		DecryptPwd string
	)
	DecryptStr, err = crypto.NewAesCBC(m.password).Decrypt(str)
	if err != nil {
		fmt.Printf("[manager] Decrypt str(%s) by key(%s) fail", str, m.password)
		return
	}
	err = json.Unmarshal([]byte(DecryptStr), &pwd)
	if err != nil {
		fmt.Printf("[manager] json.Unmarshal getStr(%s) err:%s", DecryptStr, err.Error())
		return
	}
	// Decrypt pwd again
	DecryptPwd, err = crypto.NewAesCBC(m.password).Decrypt(pwd.Pwd)
	if err != nil {
		fmt.Printf("[manager] Decrypt pwd(%s) by key(%s) fail", pwd.Pwd, m.password)
		return
	}
	pwd.Pwd = DecryptPwd
	return
}