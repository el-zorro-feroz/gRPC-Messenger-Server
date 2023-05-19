package routes

import (
	controller "main/src/api/controllers"
	"time"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(
	group *gin.RouterGroup,
	controller controller.ProfileController,
	timeout time.Duration,
) {
	group.GET("/profile")
}
