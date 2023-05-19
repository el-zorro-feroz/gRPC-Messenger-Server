package usecase

import (
	"context"
	"main/src/domain/entities"
	domain "main/src/domain/repositories"
	"main/src/utils"
	"time"
)

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string, timeout time.Duration) (entities.User, error)
	CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entities.User, secret string, expiry int) (accessToken string, err error)
}

type loginUsecase struct {
	userRepository domain.UserRepository
}

func NewLoginUsecase(userRepository domain.UserRepository) LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string, timeout time.Duration) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
