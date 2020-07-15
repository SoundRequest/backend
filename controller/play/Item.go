package play

import (
	"net/http"
	"strconv"

	"github.com/SoundRequest/backend/db"
	"github.com/SoundRequest/backend/structure/request"
	"github.com/gin-gonic/gin"
)

func NotSupport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This Function is't maked yet",
	})
}

func AddSong(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.AddSong
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if err := db.AddSong(id, body.Name, body.Description, body.Link); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Adding Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func UpdateSong(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.UpdateSong
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if err := db.UpdateSong(id, body.Target, body.Name, body.Description, body.Link); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Updating Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func RemoveSong(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	target, errCannotReturnInt := strconv.Atoi(c.DefaultQuery("target", "0"))
	if target == 0 || errCannotReturnInt != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if err := db.RemoveSong(id, target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	return
}

func GetSongs(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	result, err := db.GetSongs(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Getting Songs",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"list":    result,
	})
	return
}
