package provider

import (
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
