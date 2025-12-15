package dashboard

import (
	"github.com/futek/donation-campaign/internal/application"
	"github.com/futek/donation-campaign/internal/infrastructure/http/dashboard/handler"
	"github.com/gin-gonic/gin"
)

// @Summary Ping dashboard
// @Description Ping the dashboard
// @Tags dashboard
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /dashboard/ping [get]
func RegisterRoutes(r *gin.Engine, userService application.UserService) {
	userHandler := handler.NewUserHandler(userService)
	dashboard := r.Group("/dashboard")
	{
		dashboard.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong from dashboard",
			})
		})
	}

	userRoutes := dashboard.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
	}
}
