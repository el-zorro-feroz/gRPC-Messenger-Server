package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(timeout time.Duration, group *gin.RouterGroup) {
	group.POST("/signup")
}
