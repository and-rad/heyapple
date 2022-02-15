package app

import "golang.org/x/crypto/bcrypt"

// Crypter decrypts and encrypts passwords.
type Crypter interface {
	Encrypt(pass string) string
	Match(hash, pass string) bool
}

type crypter struct {
	cost int
}

// Encrypt returns an encrypted hash of a plaintext string.
func (c crypter) Encrypt(s string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(s), c.cost)
	if err != nil {
		panic(err)
	}

	return string(h)
}

// Match returns true if the plaintext password can be
// encrypted to produce the hash that it's tested against.
func (c crypter) Match(hash, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)) == nil
}

// NewCrypter returns a default implementation of the
// Crypter interface. The cost parameter is used for the
// cost of the underlying hash function.
func NewCrypter() Crypter {
	return crypter{
		cost: getConfig().encryptCost,
	}
}
