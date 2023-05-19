package routes

import (
	controller "main/src/api/controllers"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(
	group *gin.RouterGroup,
	controller controller.SignupController,
	timeout time.Duration,
) {
	group.POST("/signup", controller.Signup)
}
