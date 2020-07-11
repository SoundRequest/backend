package controller

import (
	"net/http"

	"github.com/SoundRequest/backend/db"
	"github.com/SoundRequest/backend/helper"
	"github.com/SoundRequest/backend/structure/request"
	"github.com/gin-gonic/gin"
)

type Auth struct {
}

func NewAuth() *Auth {
	return &Auth{}
}

// SignIn controls Sign In
func (a *Auth) SignIn(c *gin.Context) {
	var body request.SignIn
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	result, err := db.FindUser(body.Name)
	if err == db.ErrUserNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "UserNotFound",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}
	if result := helper.CheckPassword(body.Password, result.Password); result == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password NotMatch",
		})
		return
	}

	token, err := helper.GetJwtToken(result.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"token":    token,
			"verified": result.Verified,
		})
	}
}

// SignUp controller
func (a *Auth) SignUp(c *gin.Context) {
	var body request.SignUp
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if err := db.CreateUser(body.Email, body.Name, body.Password); err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": "해당 이메일 주소는 이미 가입되어 있습니다."})
		return
	}

	result, errUserNotFound := db.FindUser(body.Name)
	if errUserNotFound == db.ErrUserNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "UserNotFound",
		})
		return
	}

	token, errGetJwtToken := helper.GetJwtToken(result.ID)
	if errSendVerifyMail := helper.SendVefiryMail(result.VerifyCode, body.Email); errSendVerifyMail != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "errSendVerifyMail",
		})
		return
	}
	if errGetJwtToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "errGetJwtToken",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"verified": false,
	})
}

// Status of JWT TOKEN
func (a *Auth) Status(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	result, errUserNotFound := db.FindUserByID(id)
	if errUserNotFound != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"success": true, "verified": result.Verified, "id": id})
}

// Verify Email Handler Link
func (a *Auth) Verify(c *gin.Context) {
	code := c.Param("code")
	check, errCheckVerify := db.CheckVerify(code)
	if errCheckVerify == db.ErrUserAlreadyVerified {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "이미 인증된 계정입니다.", "success": false})
		return
	}
	if errCheckVerify != nil {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "서비스 처리중 에러가 발생하였습니다", "success": false})
		return
	}

	if check {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "정상적으로 메일인증이 되었습니다.\n이제 창을닫고 로그인하실 수 있습니다", "success": true})
	} else {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "URL을 찾지 못하였습니다.\n제대로된 요청인지 확인해주세요", "success": false})
	}
}

// UpdatePassword Controller
func (a *Auth) UpdatePassword(c *gin.Context) {
	var body request.UpdatePassword
	if errBindJSON := c.ShouldBindJSON(&body); errBindJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	result, errUserNotFound := db.FindUser(body.Name)
	if errUserNotFound == db.ErrUserNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "UserNotFound",
		})
		return
	}
	if errUserNotFound != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	if result := helper.CheckPassword(body.Origin, result.Password); result == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Origin Password NotMatch",
		})
		return
	}

	errUpdatePassword := db.UpdatePassword(result.Name, body.New)
	if errUpdatePassword != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful",
	})
}

// RecoverPasswordVerifyCode send verify code to email
func (a *Auth) RecoverPasswordVerifyCode(c *gin.Context) {
	var body request.SendVerifyPasswordCode
	if errBindJSON := c.ShouldBindJSON(&body); errBindJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	code := helper.CreateRandomString(8)
	if err := db.SetPasswordVerifyCode(body.Email, code); err != nil {
		if err == db.ErrUserNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "UserNotFound",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successed to send",
	})
}

// RecoverPassword With VerifyCode
func (a *Auth) RecoverPassword(c *gin.Context) {
	var body request.PasswordWithCode
	if errBindJSON := c.ShouldBindJSON(&body); errBindJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	okay, err := db.CheckPasswordVerify(body.Code, body.New)
	if okay != true {
		if err == db.ErrUserNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "UserNotFound",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successed to Change",
	})
}
