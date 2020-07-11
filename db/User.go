package db

import (
	"github.com/SoundRequest/backend/helper"
	"github.com/SoundRequest/backend/structure"
	"github.com/jinzhu/gorm"
)

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
	verifyCode := helper.CreateRandomString(8)

	hashedPassword := helper.HashAndSalt(password)
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

// UpdatePassword updates user's password
func UpdatePassword(name, password string) error {
	result := DB().Model(&structure.User{}).Where("name = ?", name).Update("password", helper.HashAndSalt(password))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CheckPasswordVerify handles password change verify codes
func CheckPasswordVerify(code, new string) (bool, error) {
	if code == "" {
		return false, nil
	}
	data := &structure.User{}
	result := DB().Where("verify_code_password = ?", code).First(&data)
	if result.Error == gorm.ErrRecordNotFound {
		return false, ErrUserNotFound
	} else if result.Error != nil {
		return false, result.Error
	}

	newhash := helper.HashAndSalt(new)

	if result := DB().Model(&structure.User{}).Where("verify_code_password = ?", code).UpdateColumns(structure.User{VerifyCodePassword: "", Password: newhash}); result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

// SetPasswordVerifyCode handles password change event
func SetPasswordVerifyCode(email, code string) error {
	data := &structure.User{}
	result := DB().Where("email = ?", email).First(&data)
	if result.Error == gorm.ErrRecordNotFound {
		return ErrUserNotFound
	} else if result.Error != nil {
		return result.Error
	}

	if result := DB().Model(&structure.User{}).Where("email = ?", email).Update("verify_code_password", code); result.Error != nil {
		return result.Error
	}
	return nil
}
