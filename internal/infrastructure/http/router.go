package http

import (
	"github.com/futek/donation-campaign/internal/application"
	"github.com/futek/donation-campaign/internal/infrastructure/http/dashboard"
	"github.com/futek/donation-campaign/internal/infrastructure/http/frontend"
	"github.com/gin-gonic/gin"
)

func NewRouter(userService application.UserService) *gin.Engine {
	r := gin.Default()

	dashboard.RegisterRoutes(r, userService)
	frontend.RegisterRoutes(r)

	return r
}
