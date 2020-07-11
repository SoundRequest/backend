package controller

import (
	"net/http"
	"strconv"

	"github.com/SoundRequest/backend/db"
	"github.com/SoundRequest/backend/structure/request"
	"github.com/gin-gonic/gin"
)

type Play struct{}

func NewPlay() *Play {
	return &Play{}
}

func (p *Play) NotSupport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This Function is't maked yet",
	})
}

func (p *Play) AddSong(c *gin.Context) {
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

func (p *Play) UpdateSong(c *gin.Context) {
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

func (p *Play) RemoveSong(c *gin.Context) {
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

func (p *Play) GetList(c *gin.Context) {
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

func (p *Play) GetListDetail(c *gin.Context) {
	//TODO: GetListDetail Handle
}

func (p *Play) AddList(c *gin.Context) {
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

func (p *Play) UpdateList(c *gin.Context) {
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

func (p *Play) RemoveList(c *gin.Context) {
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

func (p *Play) GetTag(c *gin.Context) {
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

func (p *Play) GetTagDetail(c *gin.Context) {
	//TODO: GetTagDetail Handle
}

func (p *Play) AddTag(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.AddTag
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})
		return
	}
	if err := db.AddTag(id, body.Name, body.Description, *body.Public); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (p *Play) UpdateTag(c *gin.Context) {
	id, _ := c.MustGet("UserId").(int)
	var body request.UpdateTag
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}
	if err := db.UpdateTag(id, body.Target, body.Name, body.Description, *body.Public); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Server Error While Removing Song",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (p *Play) RemoveTag(c *gin.Context) {
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
