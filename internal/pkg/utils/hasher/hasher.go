package hasher

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

const SaltLen = 8

func HashPass(plainPassword string) []byte {
	salt := make([]byte, SaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return nil
	} // TODO: handle error

	return hash(salt, plainPassword)
}

func CheckPass(passHash []byte, plainPassword string) bool {
	salt := make([]byte, SaltLen)
	copy(salt, passHash[:SaltLen])

	userPassHash := hash(salt, plainPassword)

	return bytes.Equal(userPassHash, passHash)
}

func hash(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)

	return append(salt, hashedPass...)
}
