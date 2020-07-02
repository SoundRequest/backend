package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt password
func HashAndSalt(pwd []byte) string {
	hash, hashAndSaltErr := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if hashAndSaltErr != nil {
		log.Println(hashAndSaltErr)
	}
	return string(hash)
}
