package routes

import (
	"github.com/SoundRequest/backend/controller"
	"github.com/SoundRequest/backend/helper/middleware"
	"github.com/gin-gonic/gin"
)

// Play Manage
func Play(route *gin.Engine) {
	play := route.Group("/play")
	c := controller.NewPlay()
	play.Use(middleware.CheckAuth){
		play.GET("/", c.NotSupport)
		play.POST("/", c.NotSupport)
		play.PATCH("/", c.NotSupport)
		play.DELETE("/", c.NotSupport)
		play.GET("/tag", c.NotSupport)
		play.GET("/list", c.NotSupport)
		play.POST("/tag", c.NotSupport)
		play.PATCH("/tag", c.NotSupport)
		play.DELETE("/tag", c.NotSupport)
		play.POST("/list", c.NotSupport)
		play.PATCH("/list", c.NotSupport)
		play.DELETE("/list", c.NotSupport)
	}
}
