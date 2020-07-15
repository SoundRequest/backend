package play

import (
	"net/http"
	"strconv"

	"github.com/SoundRequest/backend/db"
	"github.com/SoundRequest/backend/structure/request"
	"github.com/gin-gonic/gin"
)

func GetTag(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	result, err := db.GetTag(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list": result,
	})
	return
}

func GetTagDetail(c *gin.Context) {
}

func AddTag(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.AddTag
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})
		return
	}
	if err := db.AddTag(id, body.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func UpdateTag(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.UpdateTag
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if err := db.UpdateTag(id, body.Target, body.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func RemoveTag(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	target, errCannotReturnInt := strconv.Atoi(c.DefaultQuery("target", "0"))
	if target == 0 || errCannotReturnInt != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if err := db.RemoveTag(id, target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing Tag",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	return
}
