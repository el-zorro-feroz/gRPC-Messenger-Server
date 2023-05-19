package cmd

import (
	"log"
	"main/src/api/routes"
	"main/src/provider"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func RunServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("dotEnv: can't loading .env file")
	}

	app := fx.New(
		provider.UsecaseModule,
		fx.Invoke(routes.Setup),
	)
	app.Run()

	return nil
}
