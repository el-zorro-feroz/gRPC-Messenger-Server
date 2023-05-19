package provider

import (
	controller "main/src/api/controllers"
	usecase "main/src/domain/usecases"

	"go.uber.org/fx"
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
