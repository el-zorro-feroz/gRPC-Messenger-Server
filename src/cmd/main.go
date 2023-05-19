package cmd

import (
	"context"
	"log"
	"main/src/api/routes"
	"main/src/provider"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func registerHooks(lifecycle fx.Lifecycle, routes routes.Routes) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				log.Print("Starting server.")

				engine, err := routes.Setup()
				if err != nil {
					return err
				}

				serverAddr := os.Getenv("SERVER_ADDR")
				serverPort := os.Getenv("SERVER_PORT")

				engine.Run(strings.Join([]string{serverAddr, serverPort}, ":"))

				return nil
			},
			OnStop: func(context.Context) error {
				log.Print("Stopping server.")
				return nil
			},
		},
	)
}

func RunServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("dotEnv: can't loading .env file")
	}

	app := fx.New(
		fx.Provide(NewLogger),
		provider.UsecaseModule,
		provider.ControllerModule,
		fx.Provide(routes.NewRoutes),
		fx.Invoke(registerHooks),
	)
	app.Run()

	return nil
}

// NewLogger constructs a logger.
func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	return logger
}
