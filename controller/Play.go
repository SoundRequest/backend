package controller

import (
	"net/http"

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
