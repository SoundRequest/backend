package routes

import (
	"github.com/SoundRequest/backend/controller"
	"github.com/gin-gonic/gin"
)

// Auth Manage
func Auth(route *gin.Engine) {
	auth := route.Group("/auth")
	auth.POST("/signin", controller.SignIn)
}
