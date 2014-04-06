package util

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/rand"
	"github.com/evantbyrne/evanbyrne-go/config"
)

const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()_-+={}[]:;<>,./|\\?"

func ComparePassword(hashed string, password string, userSalt string) bool {
	bytesHashed := []byte(hashed)
	bytesPassword := []byte(password + userSalt + config.HashSalt)
	if err := bcrypt.CompareHashAndPassword(bytesHashed, bytesPassword); err == nil {
		return true
	}

	return false
}

func Hash(password string, userSalt string) string {
	bytes := []byte(password + userSalt + config.HashSalt)
	hash, err := bcrypt.GenerateFromPassword(bytes, config.HashCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func Random(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = characters[b % byte(len(characters))]
	}

	return string(bytes)
}