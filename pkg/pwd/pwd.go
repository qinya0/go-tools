package pwd

import "fmt"

type PWD struct {
	Key string `json:"key"`
	Pwd string `json:"pwd"`
	Msg string `json:"msg"`
}

func NewPWd(key, pwd, msg string) *PWD {
	return &PWD{
		Key: key,
		Pwd: pwd,
		Msg: msg,
	}
}


func (p *PWD) String() string {
	return fmt.Sprintf("k:%s \tp:%s \tm:%s", p.Key, p.Pwd, p.Msg)
}