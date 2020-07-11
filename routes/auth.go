package routes

import (
	"github.com/SoundRequest/backend/controller"
	"github.com/SoundRequest/backend/helper/middleware"
	"github.com/gin-gonic/gin"
)

// Auth Manage
func Auth(route *gin.Engine) {
	auth := route.Group("/auth")
	c := controller.NewAuth()
	auth.GET("/status", middleware.CheckAuth(), c.Status)
	auth.POST("/signin", c.SignIn)
	auth.POST("/signup", c.SignUp)
	auth.POST("/updatepassword", c.UpdatePassword)
	auth.POST("/recoverPasswordVerifyCode", c.RecoverPasswordVerifyCode)
	auth.POST("/recoverpassword", c.RecoverPassword)
	auth.GET("/verify/:code", c.Verify)
}
