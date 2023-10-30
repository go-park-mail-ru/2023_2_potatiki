package hasher

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

const SaltLen = 8

func HashPass(plainPassword string) []byte {
	salt := make([]byte, SaltLen)
	rand.Read(salt)

	return hash(salt, plainPassword)
}

func CheckPass(passHash []byte, plainPassword string) bool {
	salt := passHash[0:SaltLen]
	userPassHash := hash(salt, plainPassword)

	return bytes.Equal(userPassHash, passHash)
}

func hash(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), []byte(salt), 1, 64*1024, 4, 32)

	return append(salt, hashedPass...)
}
