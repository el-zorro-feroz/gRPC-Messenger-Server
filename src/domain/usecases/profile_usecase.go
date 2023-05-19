package usecase

import (
	"context"
	"main/src/domain/entities"
	domain "main/src/domain/repositories"
	"time"
)

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*entities.Profile, error)
}

type profileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*entities.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &entities.Profile{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
