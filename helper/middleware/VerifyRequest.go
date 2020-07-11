package middleware

import (
	"github.com/gin-gonic/gin"
)

func VerifyRequest(body interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var body body
		// if err := c.ShouldBindJSON(&body); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"message": "Bad Request",
		// 	})
		// 	return
		// }
	}
}
