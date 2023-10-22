package repository

import (
	"context"

	"github.com/danluki/effective-mobile-test/internal/database/models"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Users interface {
	Create(ctx context.Context, userCreateInfo CreateUserInfo) (*models.User, error)
	GetMany(ctx context.Context, input GetManyInput) ([]models.User, error)
	Delete(ctx context.Context, id int32) error
	Update(ctx context.Context, id int32, userUpdateInfo UpdateUserInfo) (*models.User, error)
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *gorm.DB) Repositories {
	return Repositories{
		Users: NewUsersRepo(db),
	}
}
