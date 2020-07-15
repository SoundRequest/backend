package routes

import (
	c "github.com/SoundRequest/backend/controller/play"
	"github.com/SoundRequest/backend/helper/middleware"
	"github.com/gin-gonic/gin"
)

// Play Manage
func Play(route *gin.Engine) {
	play := route.Group("/play")
	play.Use(middleware.CheckAuth())
	// PlayItem
	play.GET("/", c.GetSongs)
	play.POST("/", c.AddSong)
	play.PATCH("/", c.UpdateSong)
	play.DELETE("/", c.RemoveSong)
	// PlayTag
	play.GET("/tag", c.NotSupport)
	play.POST("/tag", c.NotSupport)
	play.PATCH("/tag", c.NotSupport)
	play.DELETE("/tag", c.NotSupport)
	// PlayList
	play.GET("/list", c.GetList)
	play.GET("/list/detail", c.GetListDetail)
	play.POST("/list", c.AddList)
	play.PATCH("/list", c.UpdateList)
	play.DELETE("/list", c.RemoveList)
	// Bridge (Connect to Playlist OR Tag)
	play.POST("/bridge/tag", c.AddTagToItem)
	play.POST("/bridge/list", c.AddPlayListToItem)
	play.DELETE("/bridge/tag", c.RemoveTagFromItem)
	play.DELETE("/bridge/list", c.RemovePlayListFromItem)

}
