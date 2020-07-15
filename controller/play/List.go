package play

import (
	"net/http"
	"strconv"

	"github.com/SoundRequest/backend/db"
	"github.com/SoundRequest/backend/structure/request"
	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	result, err := db.GetList(id)
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

func GetListDetail(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	target, errCannotReturnInt := strconv.Atoi(c.DefaultQuery("target", "0"))
	if target == 0 || errCannotReturnInt != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	result, err := db.GetListDetail(id, target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Get Detail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list": result,
	})
	return
}

func AddList(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.AddList
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})
		return
	}
	if err := db.AddList(id, body.Name, body.Description, *body.Public); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func UpdateList(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.UpdateList
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if err := db.UpdateList(id, body.Target, body.Name, body.Description, *body.Public); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func RemoveList(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	target, errCannotReturnInt := strconv.Atoi(c.DefaultQuery("target", "0"))
	if target == 0 || errCannotReturnInt != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if err := db.RemoveList(id, target); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	return
}
