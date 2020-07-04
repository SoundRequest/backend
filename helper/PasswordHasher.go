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

// CheckPassowrd check password
func CheckPassowrd(normal, hashed string) bool {

	verifyHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(verifyHash, []byte(normal))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
