package routes

import (
	c "github.com/SoundRequest/backend/controller/chart"
	"github.com/gin-gonic/gin"
)

// Chart Manage
func Chart(route *gin.Engine) {
	chart := route.Group("/chart")

	chart.GET("/melon", c.Melon)
}
