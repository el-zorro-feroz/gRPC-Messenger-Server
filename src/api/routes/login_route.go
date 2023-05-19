package routes

import (
	controller "main/src/api/controllers"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(
	group *gin.RouterGroup,
	controller controller.LoginController,
	timeout time.Duration,
) {
	group.POST("/login", controller.Login)
}
