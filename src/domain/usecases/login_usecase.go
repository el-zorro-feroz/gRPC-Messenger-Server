package usecase

import (
	"context"
	"main/src/domain/entities"
	domain "main/src/domain/repositories"
	"main/src/utils"
	"time"
)

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (entities.User, error)
	CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entities.User, secret string, expiry int) (accessToken string, err error)
}

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
