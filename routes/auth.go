package routes

import (
	c "github.com/SoundRequest/backend/controller/auth"
	"github.com/SoundRequest/backend/helper/middleware"
	"github.com/gin-gonic/gin"
)

// Auth Manage
func Auth(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.GET("/status", middleware.CheckAuth(), c.Status)
	auth.POST("/signin", c.SignIn)
	auth.POST("/signup", c.SignUp)
	auth.GET("/profile", middleware.CheckAuth(), c.Profile)
	auth.POST("/updatepassword", c.UpdatePassword)
	auth.POST("/recoverPasswordVerifyCode", c.RecoverPasswordVerifyCode)
	auth.POST("/recoverpassword", c.RecoverPassword)
	auth.GET("/verify/:code", c.Verify)
}
