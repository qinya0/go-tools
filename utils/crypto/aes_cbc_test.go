package crypto

import (
	"fmt"
	"testing"
)

var (
	testKey = "key"
	testAes = Aes{Key: testKey}
)

func TestAes(t *testing.T) {
	// encode
	value := "testValue"
	enStr :=  testAes.Encrypt(value)
	fmt.Println("enStr:", enStr)

	// decode
	deStr :=  testAes.Decrypt(enStr)
	fmt.Println("deStr:", deStr)
	if value != deStr {
		t.Errorf("value(%s) not eq deStr(%s)", value, deStr)
	}
}
