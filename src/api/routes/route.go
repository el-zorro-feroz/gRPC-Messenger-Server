package routes

import (
	controller "main/src/api/controllers"
	"main/src/api/middleware"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Routes interface {
	Setup() (*gin.Engine, error)
}

type routes struct {
	signupController       controller.SignupController
	loginController        controller.LoginController
	refreshTokenController controller.RefreshTokenController
	profileController      controller.ProfileController
}

func NewRoutes(
	signupController controller.SignupController,
	loginController controller.LoginController,
	refreshTokenController controller.RefreshTokenController,
	profileController controller.ProfileController,
) Routes {
	return &routes{
		signupController:       signupController,
		loginController:        loginController,
		refreshTokenController: refreshTokenController,
		profileController:      profileController,
	}
}

func (r routes) Setup() (*gin.Engine, error) {
	engine := gin.Default()
	timeout := time.Duration(2) * time.Second

	publicRouter := engine.Group("")

	NewSignupRouter(publicRouter, r.signupController, timeout)
	NewLoginRouter(publicRouter, r.loginController, timeout)
	NewRefreshTokenRouter(publicRouter, r.refreshTokenController, timeout)

	accessTokenSecret := os.Getenv("SERVER_SECRET")

	protectedRouter := engine.Group("")
	protectedRouter.Use(middleware.JwtMiddleware(accessTokenSecret))

	NewProfileRouter(protectedRouter, r.profileController, timeout)

	return engine, nil
}
