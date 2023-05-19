package provider

import (
	controller "main/src/api/controllers"
	domain "main/src/domain/repositories"
	usecase "main/src/domain/usecases"

	"go.uber.org/fx"
)

// UserRepository .
var UserRepository = fx.Options(
	fx.Provide(domain.NewSqliteUserRepository),
)

// UsecaseModule .
var UsecaseModule = fx.Options(
	fx.Provide(usecase.NewLoginUsecase),
	fx.Provide(usecase.NewSignupUsecase),
	fx.Provide(usecase.NewProfileUsecase),
	fx.Provide(usecase.NewRefreshTokenUsecase),
)

// ControllerModule .
var ControllerModule = fx.Options(
	fx.Provide(controller.NewLoginController),
	fx.Provide(controller.NewSignupController),
	fx.Provide(controller.NewProfileController),
	fx.Provide(controller.NewRefreshTokenUsecase),
)
