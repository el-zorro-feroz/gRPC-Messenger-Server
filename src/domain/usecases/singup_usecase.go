package usecase

import (
	"context"
	"main/src/domain/entities"
	domain "main/src/domain/repositories"
	"main/src/utils"
	"time"
)

type SignupUsecase interface {
	Signup(c context.Context, user *entities.User, timeout time.Duration) error
	GetUserByEmail(c context.Context, email string, timeout time.Duration) (entities.User, error)
	CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error)
}

type signupUsecase struct {
	userRepository domain.UserRepository
}

func NewSignupUsecase(userRepository domain.UserRepository) SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
	}
}

func (su *signupUsecase) Signup(c context.Context, user *entities.User, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string, timeout time.Duration) (entities.User, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, secret, expiry)
}
