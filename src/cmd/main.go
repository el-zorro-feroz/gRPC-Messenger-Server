package cmd

import (
	"log"
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
	)
	app.Run()

	return nil
}
