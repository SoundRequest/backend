package db

import (
	"math/rand"
	"strings"

	"github.com/SoundRequest/backend/structure"
	"github.com/jinzhu/gorm"
)

var chars []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" +
	"0123456789")

// FindUserByID controller
func FindUserByID(id string) (*structure.User, error) {
	return &structure.User{}, nil
}

// FindUserByEmail controller
func FindUserByEmail(email string) (*structure.User, error) {
	return &structure.User{}, nil
}

// FindUserByName controller
func FindUserByName(name string) (*structure.User, error) {
	return &structure.User{}, nil
}

// CreateUser controller
func CreateUser(email, name, password string) error {
	// generate Random Char
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	verifyCode := b.String()

	passwordHash := password
	result := DB().Create(&structure.User{Name: name, Email: email, Password: passwordHash, VerifyCode: verifyCode})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CheckVerify For Email Check
func CheckVerify(code string) (bool, error) {
	data := &structure.User{}
	result := DB().Where("verify_code = ?", code).First(&data)
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	if data.Verified {
		return false, ErrUserAlreadyVerified
	}

	_result := DB().Model(&structure.User{}).Where("verify_code = ?", code).Update("verified", true)
	if _result.Error != nil {
		return false, _result.Error
	}
	return true, nil
}
