package util

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/evantbyrne/evanbyrne-go/config"
)

func Hash(password string, userSalt string) string {
	bytes := []byte(password + userSalt + config.HashSalt)
	hash, err := bcrypt.GenerateFromPassword(bytes, config.HashCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}