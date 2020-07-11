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
	play.Use(middleware.CheckAuth())
	play.GET("/", c.NotSupport)
	play.POST("/", c.AddSong)
	play.PATCH("/", c.UpdateSong)
	play.DELETE("/", c.RemoveSong)
	play.GET("/tag", c.NotSupport)
	play.POST("/tag", c.NotSupport)
	play.PATCH("/tag", c.NotSupport)
	play.DELETE("/tag", c.NotSupport)
	play.GET("/list", c.GetList)
	play.GET("/list/detail", c.GetListDetail)
	play.POST("/list", c.AddList)
	play.PATCH("/list", c.UpdateList)
	play.DELETE("/list", c.RemoveList)
}
