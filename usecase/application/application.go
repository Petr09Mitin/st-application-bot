package application

import (
	"context"

	"github.com/st-application-bot/domain"
)

type ApplicationStorage interface {
	Store(ctx context.Context, application *domain.Application) (newApplication *domain.Application, err error)
	GetAll(ctx context.Context) (applications []*domain.Application, err error)
}

type Service struct {
	storage ApplicationStorage
}

func NewService(storage ApplicationStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateApplication(ctx context.Context, application *domain.Application) (newApplication *domain.Application, err error) {
	newApplication, err = s.storage.Store(ctx, application)
	if err != nil {
		return
	}

	return
}

func (s *Service) GetAllApplications(ctx context.Context) (applications []*domain.Application, err error) {
	applications, err = s.storage.GetAll(ctx)
	if err != nil {
		return
	}

	return
}
