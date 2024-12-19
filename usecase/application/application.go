package application

import (
	"context"

	"github.com/st-application-bot/domain"
)

type ApplicationStorage interface {
	Store(ctx context.Context, application *domain.Application) (newApplication *domain.Application, err error)
	GetAll(ctx context.Context) (applications []*domain.Application, err error)
	GetUnresolved(ctx context.Context) (applications []*domain.Application, err error)
	GetApproved(ctx context.Context) (applications []*domain.Application, err error)
	GetRejected(ctx context.Context) (applications []*domain.Application, err error)
	GetByTgUserID(ctx context.Context, tgUserID int64) (application *domain.Application, err error)
	GetByID(ctx context.Context, applicationID uint) (application *domain.Application, err error)
	Update(ctx context.Context, application *domain.Application) (updatedApplication *domain.Application, err error)
}

type Service struct {
	storage ApplicationStorage
}

func NewService(storage ApplicationStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetApplicationByTgUserID(ctx context.Context, tgUserID int64) (application *domain.Application, err error) {
	application, err = s.storage.GetByTgUserID(ctx, tgUserID)
	if err != nil {
		return
	}

	return
}

func (s *Service) CreateApplication(ctx context.Context, application *domain.Application) (newApplication *domain.Application, err error) {
	_, err = s.storage.GetByTgUserID(ctx, application.TelegramUserID)
	if err != domain.ErrApplicationNotFound {
		err = domain.ErrApplicationAlreadyExists
		return
	}

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

func (s *Service) GetUnresolvedApplications(ctx context.Context) (applications []*domain.Application, err error) {
	applications, err = s.storage.GetUnresolved(ctx)
	if err != nil {
		return
	}

	return
}

func (s *Service) ApproveApplication(ctx context.Context, applicationID uint) (application *domain.Application, err error) {
	application, err = s.storage.GetByID(ctx, applicationID)
	if err != nil {
		return
	}

	application.Approved = true
	application.Resolved = true

	application, err = s.storage.Update(ctx, application)
	if err != nil {
		return
	}

	return
}

func (s *Service) RejectApplication(ctx context.Context, applicationID uint) (application *domain.Application, err error) {
	application, err = s.storage.GetByID(ctx, applicationID)
	if err != nil {
		return
	}

	application.Approved = false
	application.Resolved = true

	application, err = s.storage.Update(ctx, application)
	if err != nil {
		return
	}

	return
}

func (s *Service) GetApprovedApplications(ctx context.Context) (applications []*domain.Application, err error) {
	applications, err = s.storage.GetApproved(ctx)
	if err != nil {
		return
	}

	return
}

func (s *Service) GetRejectedApplications(ctx context.Context) (applications []*domain.Application, err error) {
	applications, err = s.storage.GetRejected(ctx)
	if err != nil {
		return
	}

	return
}
