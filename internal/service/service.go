package service

import (
	"context"

	"github.com/SteelPangolin/go-genderize"
	"github.com/danluki/effective-mobile-test/internal/domain"
	"github.com/danluki/effective-mobile-test/internal/logger"
	"github.com/danluki/effective-mobile-test/internal/repository"
	"github.com/masonkmeyer/agify"
	"github.com/masonkmeyer/nationalize"
)

//go:generate mockgen -source service.go -destination=mocks/mock.go

type Users interface {
	Create(ctx context.Context, input CreateUserInput) (*domain.User, error)
	List(ctx context.Context, input ListUsersInput) ([]domain.User, error)
	Update(ctx context.Context, input UpdateUserInput) (*domain.User, error)
	Delete(ctx context.Context, id int32) error
}

type Services struct {
	Users Users
}

type GenderizeClientInterface interface {
	Get(query genderize.Query) ([]genderize.Response, error)
}

type NationalizeClientInterface interface {
	Predict(name string) (*nationalize.Prediction, *nationalize.RateLimit, error)
}

type AgifyClientInterface interface {
	Predict(name string) (*agify.Prediction, *agify.RateLimit, error)
}

type Deps struct {
	Repositories      repository.Repositories
	GenderizeClient   GenderizeClientInterface
	NationalizeClient NationalizeClientInterface
	AgifyClient       AgifyClientInterface
	Logger            logger.Logger
}

func NewServices(deps Deps) *Services {
	return &Services{
		Users: NewUsersService(
			deps.Repositories,
			deps.GenderizeClient,
			deps.NationalizeClient,
			deps.AgifyClient,
			deps.Logger,
		),
	}
}
