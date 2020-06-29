package controller

import "github.com/gin-gonic/gin"

// SignIn controlls Sign In
func SignIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": "yes",
	})
}
