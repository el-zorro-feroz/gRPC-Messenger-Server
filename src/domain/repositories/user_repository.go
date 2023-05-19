package domain

import (
	"context"
	"main/src/domain/entities"
)

type UserRepository interface {
	GetByEmail(context.Context, string) (entities.User, error)
}
