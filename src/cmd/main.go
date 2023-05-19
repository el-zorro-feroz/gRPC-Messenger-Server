package cmd

import (
	"context"
	"log"
	"main/src/api/routes"
	"main/src/provider"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func registerHooks(lifecycle fx.Lifecycle) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				log.Print("Starting server.")

				timeout := time.Duration(2) * time.Second
				gin := gin.Default()

				if err := routes.Setup(timeout, gin); err != nil {
					return err
				}

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
