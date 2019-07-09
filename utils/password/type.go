package password

import (
	"fmt"
	"sync"
)

var (
	allStr      = "abcdefghijklmnopqrstuvwxyz"
	allUpperStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	allInt      = "0123456789"
	allSpecial  = "!@#$%^&*"
	allPointer  = "_-:;,."

	defaultManager = TypeManager{
		Types: map[string]Getter{},
		Mutex: sync.Mutex{},
	}
)

func init() {
	// registry all type
	RegistryType(&typeStr{})
	RegistryType(&typeUpperStr{})
	RegistryType(&typeInt{})
	RegistryType(&typeSpecial{})
	RegistryType(&typePointer{})
}

type Getter interface{
	GetName() string
	GetAllValue() string
}

type typeStr struct {}
type typeUpperStr struct {}
type typeInt struct {}
type typeSpecial struct {}
type typePointer struct {}

func (t *typeStr) GetName() string {
	return "typeStr"
}
func (t *typeStr) GetAllValue() string {
	return allStr
}
func (t *typeUpperStr) GetName() string {
	return "typeUpperStr"
}
func (t *typeUpperStr) GetAllValue() string {
	return allUpperStr
}
func (t *typeInt) GetName() string {
	return "typeInt"
}
func (t *typeInt) GetAllValue() string {
	return allInt
}
func (t *typeSpecial) GetName() string {
	return "typeSpecial"
}
func (t *typeSpecial) GetAllValue() string {
	return allSpecial
}
func (t *typePointer) GetName() string {
	return "typePointer"
}
func (t *typePointer) GetAllValue() string {
	return allPointer
}

type TypeManager struct {
	Types map[string]Getter
	sync.Mutex
}

func RegistryType(g Getter) (err error) {
	defaultManager.Lock()
	if _, ok := defaultManager.Types[g.GetName()]; !ok {
		// add
		defaultManager.Types[g.GetName()] = g
	} else {
		err = fmt.Errorf("type:%s is exist!", g.GetName())
	}
	defaultManager.Unlock()
	return
}

func DeleteType(g Getter) {
	defaultManager.Lock()
	if _, ok := defaultManager.Types[g.GetName()]; ok {
		// add
		delete(defaultManager.Types, g.GetName())
	}
	defaultManager.Unlock()
	return
}

func GetCharacter(gs []Getter) (str string) {
	defaultManager.Lock()
	for _, g := range gs{
		if _, ok := defaultManager.Types[g.GetName()]; ok {
			str += g.GetAllValue()
		}
	}
	defaultManager.Unlock()
	return
}
