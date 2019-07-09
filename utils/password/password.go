package password

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneratePassword(isUpperStr, isInt, isSpecial, isPointer bool, num int) (password string) {
	var (
		baseStr string
		Getters  []Getter
	)
	// default use string only
	Getters = append(Getters, &typeStr{})

	if isUpperStr {
		Getters = append(Getters, &typeUpperStr{})
	}
	if isInt {
		Getters = append(Getters, &typeInt{})
	}
	if isSpecial {
		Getters = append(Getters, &typeSpecial{})
	}
	if isPointer {
		Getters = append(Getters, &typePointer{})
	}

	baseStr = GetCharacter(Getters)
	password = generatePassword(num, baseStr)

	return
}

func generatePassword(number int, baseStr string) (password string) {
	var (
		i = 0
		strLen = len(baseStr)
	)
	if strLen == 0 {
		fmt.Println("[pwd] baseStr is nil!")
		return
	}
	for ; i < number; i++ {
		rand.Seed(time.Now().UTC().UnixNano() + int64(i))
		r := rand.Intn(strLen)
		password += baseStr[r:r+1]
	}
	return
}
