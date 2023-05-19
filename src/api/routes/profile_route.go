package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(timeout time.Duration, group *gin.RouterGroup) {
	group.GET("/profile")
}
