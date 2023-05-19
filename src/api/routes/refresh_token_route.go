package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(timeout time.Duration, group *gin.RouterGroup) {
	group.POST("/refresh")
}
