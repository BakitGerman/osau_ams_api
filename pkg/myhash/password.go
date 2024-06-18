package myhash

import (
	"golang.org/x/crypto/bcrypt"
)

// hash provides ...
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashPass, password string) bool
}

type Hasher struct {
	salt string
}

func NewHasher(salt string) *Hasher {
	return &Hasher{
		salt: salt,
	}
}

// hash ...
func (h *Hasher) HashPassword(password string) (string, error) {

	passWithSalt := password + h.salt
	b, err := bcrypt.GenerateFromPassword([]byte(passWithSalt), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// compare passwords ...
func (h *Hasher) ComparePassword(hashPass, password string) bool {
	passWithSalt := password + h.salt
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(passWithSalt)) == nil
}
