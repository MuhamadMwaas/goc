package frontend

import "github.com/gin-gonic/gin"

// @Summary Ping frontend
// @Description Ping the frontend
// @Tags frontend
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /frontend/ping [get]
func RegisterRoutes(r *gin.Engine) {
	frontend := r.Group("/frontend")
	{
		frontend.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong from frontend",
			})
		})
	}
}
