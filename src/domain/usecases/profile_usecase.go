package usecase

import (
	"context"
	"main/src/domain/entities"
	domain "main/src/domain/repositories"
	"time"
)

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string, timeout time.Duration) (*entities.Profile, error)
}

type profileUsecase struct {
	userRepository domain.UserRepository
}

func NewProfileUsecase(userRepository domain.UserRepository) ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string, timeout time.Duration) (*entities.Profile, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
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
