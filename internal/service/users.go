package service

import (
	"context"
	"errors"
	"log/slog"

	"github.com/SteelPangolin/go-genderize"
	"github.com/danluki/effective-mobile-test/internal/database/models"
	"github.com/danluki/effective-mobile-test/internal/domain"
	"github.com/danluki/effective-mobile-test/internal/logger"
	"github.com/danluki/effective-mobile-test/internal/repository"
)

type UsersService struct {
	repo              repository.Repositories
	genderizeClient   GenderizeClientInterface
	agifyClient       AgifyClientInterface
	nationalizeClient NationalizeClientInterface
	logger            logger.Logger
}

func NewUsersService(
	repo repository.Repositories,
	genderizeClient GenderizeClientInterface,
	nationalizeClient NationalizeClientInterface,
	agifyClient AgifyClientInterface,
	logger logger.Logger,
) *UsersService {
	return &UsersService{
		repo:              repo,
		genderizeClient:   genderizeClient,
		nationalizeClient: nationalizeClient,
		agifyClient:       agifyClient,
		logger:            logger,
	}
}

func (us *UsersService) convertUserToDomain(user *models.User) *domain.User {
	return &domain.User{
		ID:         user.ID,
		Name:       user.Name,
		Surname:    user.Surname,
		Patronymic: user.Patronymic,
		Age:        user.Age,
		Gender:     user.Gender,
		Country:    user.Country,
	}
}

type CreateUserInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func (us *UsersService) Create(ctx context.Context, input CreateUserInput) (*domain.User, error) {
	genderizeResponse, err := us.genderizeClient.Get(
		genderize.Query{Names: []string{input.Name, input.Surname, input.Surname}},
	)
	if err != nil {
		us.logger.Error("Error:", slog.Any("err", err))
		return nil, err
	}

	if len(genderizeResponse) == 0 {
		us.logger.Error("Error:", slog.Any("err", err))
		return nil, errors.New("cannot find gender")
	}

	nationalizeResponse, _, err := us.nationalizeClient.Predict(input.Name)
	if err != nil {
		us.logger.Error("Error:", slog.Any("err", err))
		return nil, err
	}

	if len(nationalizeResponse.Country) == 0 {
		us.logger.Error("Error:", slog.Any("err", err))
		return nil, errors.New("cannot find country")
	}

	agifyResponse, _, err := us.agifyClient.Predict(input.Name)
	if err != nil {
		us.logger.Error("Error:", slog.Any("err", err))
		return nil, err
	}

	user, err := us.repo.Users.Create(ctx, repository.CreateUserInfo{
		Name:       input.Name,
		Surname:    input.Surname,
		Patronymic: input.Patronymic,
		Age:        uint(agifyResponse.Age),
		Gender:     genderizeResponse[0].Gender,
		Country:    nationalizeResponse.Country[0].CountryId,
	})
	if err != nil {
		us.logger.Error("Error:", slog.Any("err", err))
		return nil, err
	}

	return us.convertUserToDomain(user), nil
}

type ListUsersInput struct {
	Gender   *string `json:"gender"`
	MinAge   *int    `json:"min_age"`
	MaxAge   *int    `json:"max_age"`
	Country  *string `json:"country"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
}

func (us *UsersService) List(ctx context.Context, input ListUsersInput) ([]domain.User, error) {
	users, err := us.repo.Users.GetMany(ctx, repository.GetManyInput{
		Gender:      input.Gender,
		MinAge:      input.MinAge,
		MaxAge:      input.MaxAge,
		Nationality: input.Country,
		Page:        input.Page,
		PageSize:    input.PageSize,
	})
	if err != nil {
		us.logger.Error("Error:", slog.Any("err", err))
		return nil, err
	}

	convertedUsers := []domain.User{}
	for _, user := range users {
		convertedUsers = append(convertedUsers, *us.convertUserToDomain(&user))
	}

	return convertedUsers, nil
}

type UpdateUserInput struct {
	ID         int     `json:"id"`
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic"`
	Age        *uint   `json:"age"`
	Gender     *string `json:"gender"`
	Country    *string `json:"country"`
}

func (us *UsersService) Update(ctx context.Context, input UpdateUserInput) (*domain.User, error) {
	user, err := us.repo.Users.Update(ctx, int32(input.ID), repository.UpdateUserInfo{
		Name:    input.Name,
		Age:     input.Age,
		Gender:  input.Gender,
		Country: input.Country,
	})
	if err != nil {
		us.logger.Error("Error:", slog.Any("err", err))
		return nil, err
	}

	return us.convertUserToDomain(user), nil
}

func (us *UsersService) Delete(ctx context.Context, id int32) error {
	err := us.repo.Users.Delete(ctx, id)
	if err != nil {
		us.logger.Error("Error:", slog.Any("err", err))
		return err
	}

	return nil
}
