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

// CheckPassword check password
func CheckPassword(normal, hashed string) bool {
	verifyHash := []byte(hashed)
	errFailedToCompareHAshAndPassword := bcrypt.CompareHashAndPassword(verifyHash, []byte(normal))
	if errFailedToCompareHAshAndPassword != nil {
		log.Println(errFailedToCompareHAshAndPassword)
		return false
	}
	return true
}
