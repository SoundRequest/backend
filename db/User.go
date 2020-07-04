package db

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/SoundRequest/backend/helper"
	"github.com/SoundRequest/backend/structure"
	"github.com/jinzhu/gorm"
)

var chars []rune = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" +
	"0123456789")

// FindUserByID controller
func FindUserByID(id int) (*structure.User, error) {
	data := &structure.User{}
	result := DB().Where("id = ?", id).First(&data)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrUserNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// FindUser With Name Or Email
func FindUser(name string) (*structure.User, error) {
	data := &structure.User{}
	result := DB().Where("name = ? OR email = ?", name, name).First(&data)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrUserNotFound
	} else if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
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

	hashedPassword := helper.HashAndSalt([]byte(password))
	fmt.Println(hashedPassword)
	result := DB().Create(&structure.User{Name: name, Email: email, Password: hashedPassword, VerifyCode: verifyCode})
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
