package routes

import (
	"github.com/SoundRequest/backend/controller"
	"github.com/SoundRequest/backend/helper/middleware"
	"github.com/gin-gonic/gin"
)

// Auth Manage
func Auth(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.GET("/status", middleware.CheckAuth(), controller.Status)
	auth.POST("/signin", controller.SignIn)
	auth.POST("/signup", controller.SignUp)
	auth.POST("/updatepassword", controller.UpdatePassword)
	auth.POST("/recoverPasswordVerifyCode", controller.RecoverPasswordVerifyCode)
	auth.POST("/recoverpassword", controller.RecoverPassword)
	auth.GET("/verify/:code", controller.Verify)
}
