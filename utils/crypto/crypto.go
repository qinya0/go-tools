package crypto

type Crypto interface {
	Encrypt(src string) (string, error)
	Decrypt(src string) (string, error)
}

