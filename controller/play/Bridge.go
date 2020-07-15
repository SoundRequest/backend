package play

import (
	"net/http"

	"github.com/SoundRequest/backend/db"
	"github.com/SoundRequest/backend/structure/request"
	"github.com/gin-gonic/gin"
)

func Temp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func AddTagToItem(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.Bridge
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})
		return
	}
	if err := db.AddTagToItem(id, body.Item, body.Target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func AddPlayListToItem(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.Bridge
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})
		return
	}
	if err := db.AddPlayListToItem(id, body.Item, body.Target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func RemoveTagFromItem(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.Bridge
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})
		return
	}
	if err := db.RemoveTagFromItem(id, body.Item, body.Target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func RemovePlayListFromItem(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.Bridge
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})
		return
	}
	if err := db.RemovePlayListFromItem(id, body.Item, body.Target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
