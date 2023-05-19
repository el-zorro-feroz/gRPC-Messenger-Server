package routes

import (
	controller "main/src/api/controllers"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(
	group *gin.RouterGroup,
	controller controller.RefreshTokenController,
	timeout time.Duration,
) {
	group.POST("/refresh", controller.RefreshToken)
}
