package domain

import (
	"context"
	"main/src/domain/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(c context.Context, user *entities.User) error
	Fetch(c context.Context) ([]entities.User, error)
	GetByEmail(c context.Context, email string) (entities.User, error)
	GetByID(c context.Context, id string) (entities.User, error)
}

type sqliteUserRepository struct {
	database *gorm.DB
}

func NewSqliteUserRepository() UserRepository {
	db, _ := gorm.Open(sqlite.Open("database.sqlite3"), &gorm.Config{})

	db.AutoMigrate(&entities.User{})

	return &sqliteUserRepository{
		database: db,
	}
}

func (sur sqliteUserRepository) Create(c context.Context, user *entities.User) error {
	tx := sur.database.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}

	sur.database.Commit()

	return nil
}

func (sur sqliteUserRepository) Fetch(c context.Context) ([]entities.User, error) {
	var users []entities.User

	tx := sur.database.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (sur sqliteUserRepository) GetByEmail(c context.Context, email string) (entities.User, error) {
	var user entities.User

	tx := sur.database.Where(&entities.User{Email: email}).First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}

	return user, nil
}

func (sur sqliteUserRepository) GetByID(c context.Context, id string) (entities.User, error) {
	var user entities.User

	tx := sur.database.Where(&entities.User{ID: id}).First(&user)
	if tx.Error != nil {
		return user, tx.Error
	}

	return user, nil
}
