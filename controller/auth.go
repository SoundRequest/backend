package controller

import (
	"net/http"

	"github.com/SoundRequest/backend/db"
	"github.com/SoundRequest/backend/helper"
	"github.com/SoundRequest/backend/structure"
	"github.com/gin-gonic/gin"
)

// SignIn controlls Sign In
func SignIn(c *gin.Context) {
	token, err := helper.GetJwtToken(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

// SignUp controller
func SignUp(c *gin.Context) {
	var body structure.SignUp
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if err := db.CreateUser(body.Email, body.Name, body.Password); err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": "해당 이메일주소는 이미 가입되어 있습니다."})
		return
	}

	token, err := helper.GetJwtToken(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

// Status of JWT TOKEN
func Status(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"success": true})
}

// Verify Email Handler Link
func Verify(c *gin.Context) {
	code := c.Param("code")
	check, err := db.CheckVerify(code)
	if err == db.ErrUserAlreadyVerified {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "이미 인증된 계정입니다.", "success": false})
		return
	}
	if err != nil {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "서비스 처리중 에러가 발생하였습니다", "success": false})
		return
	}

	if check {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "정상적으로 메일인증이 되었습니다.\n이제 창을닫고 로그인하실 수 있습니다", "success": true})
	} else {
		c.HTML(http.StatusOK, "verify.html", gin.H{"text": "URL을 찾지 못하였습니다.\n제대로된 요청인지 확인해주세요", "success": false})
	}
}
