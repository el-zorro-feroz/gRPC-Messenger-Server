package usecase

import (
	"context"
	"main/src/domain/entities"
	domain "main/src/domain/repositories"
	"main/src/utils"
	"time"
)

type RefreshTokenUsecase interface {
	GetUserByID(c context.Context, email string, timeout time.Duration) (entities.User, error)
	CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository) RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, email string, timeout time.Duration) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	return rtu.userRepository.GetByID(ctx, email)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return utils.ExtractIDFromToken(requestToken, secret)
}
