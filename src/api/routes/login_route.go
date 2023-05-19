package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(timeout time.Duration, group *gin.RouterGroup) {
	group.POST("/login")
}
