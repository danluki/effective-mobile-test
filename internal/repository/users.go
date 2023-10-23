package repository

import (
	"context"

	"github.com/danluki/effective-mobile-test/internal/database/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

type CreateUserInfo struct {
	Name       string
	Surname    string
	Patronymic string
	Age        uint
	Gender     string
	Country    string
}

func (repo *UsersRepository) Create(
	ctx context.Context,
	userCreateInfo CreateUserInfo,
) (*models.User, error) {
	user := models.User{
		Name:       userCreateInfo.Name,
		Surname:    userCreateInfo.Surname,
		Patronymic: userCreateInfo.Patronymic,
		Age:        userCreateInfo.Age,
		Gender:     userCreateInfo.Gender,
		Country:    userCreateInfo.Country,
	}

	err := repo.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type GetManyInput struct {
	Gender      *string
	MinAge      *int
	MaxAge      *int
	Nationality *string
	Page        int
	PageSize    int
}

func (repo *UsersRepository) GetMany(
	ctx context.Context,
	input GetManyInput,
) ([]models.User, error) {
	var users []models.User

	query := repo.db.Model(&models.User{})

	if input.Gender != nil {
		query = query.Where("gender = ?", *input.Gender)
	}

	if input.MinAge != nil {
		query = query.Where("age >= ?", *input.MinAge)
	}

	if input.MaxAge != nil {
		query = query.Where("age <= ?", *input.MaxAge)
	}

	if input.Nationality != nil {
		query = query.Where("nationality = ?", *input.Nationality)
	}

	offset := (input.Page - 1) * input.PageSize
	query = query.Offset(offset).Limit(input.PageSize)

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UsersRepository) Delete(ctx context.Context, id int32) error {
	err := repo.db.WithContext(ctx).Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

type UpdateUserInfo struct {
	Name    *string
	Age     *uint
	Gender  *string
	Country *string
}

func (repo *UsersRepository) Update(
	ctx context.Context,
	id int32,
	userUpdateInfo UpdateUserInfo,
) (*models.User, error) {
	var user models.User
	user.ID = id

	err := repo.db.First(&user).Error
	if err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})
	if userUpdateInfo.Name != nil {
		updates["name"] = *userUpdateInfo.Name
	}

	if userUpdateInfo.Age != nil {
		updates["age"] = *userUpdateInfo.Age
	}

	if userUpdateInfo.Gender != nil {
		updates["gender"] = *userUpdateInfo.Gender
	}

	if userUpdateInfo.Country != nil {
		updates["country"] = *userUpdateInfo.Country
	}

	if err := repo.db.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
