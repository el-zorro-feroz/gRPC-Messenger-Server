package routes

import (
	"main/src/api/middleware"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(timeout time.Duration, gin *gin.Engine) error {
	publicRouter := gin.Group("")

	NewSignupRouter(timeout, publicRouter)
	NewLoginRouter(timeout, publicRouter)
	NewRefreshTokenRouter(timeout, publicRouter)

	accessTokenSecret := os.Getenv("SERVER_SECRET")

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtMiddleware(accessTokenSecret))

	NewProfileRouter(timeout, protectedRouter)

	return nil
}
