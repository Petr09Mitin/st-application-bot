package application

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/st-application-bot/domain"
	mock_application "github.com/st-application-bot/mocks/usecase/application"
	"github.com/stretchr/testify/assert"
)

func TestApplicationService_GetApplicationByTgUserID(t *testing.T) {
	type fields struct {
		storage *mock_application.MockApplicationStorage
	}
	type args struct {
		ctx      context.Context
		tgUserID int64
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		expected    *domain.Application
		expectedErr error
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByTgUserID(gomock.Any(), gomock.Any()).Return(&domain.Application{}, nil)
			},
			args: args{
				ctx:      context.Background(),
				tgUserID: 1,
			},
			expected:    &domain.Application{},
			expectedErr: nil,
		},
		{
			name: "error not found",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByTgUserID(gomock.Any(), gomock.Any()).Return(nil, domain.ErrApplicationNotFound)
			},
			args: args{
				ctx:      context.Background(),
				tgUserID: 1,
			},
			expected:    nil,
			expectedErr: domain.ErrApplicationNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				storage: mock_application.NewMockApplicationStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := NewService(f.storage)
			got, err := s.GetApplicationByTgUserID(tt.args.ctx, tt.args.tgUserID)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestApplicationService_CreateApplication(t *testing.T) {
	type fields struct {
		storage *mock_application.MockApplicationStorage
	}
	type args struct {
		ctx  context.Context
		appl *domain.Application
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		expected    *domain.Application
		expectedErr error
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByTgUserID(gomock.Any(), gomock.Any()).Return(nil, domain.ErrApplicationNotFound)
				f.storage.EXPECT().Store(gomock.Any(), gomock.Any()).Return(&domain.Application{}, nil)
			},
			args: args{
				ctx:  context.Background(),
				appl: &domain.Application{},
			},
			expected:    &domain.Application{},
			expectedErr: nil,
		},
		{
			name: "application already exists",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByTgUserID(gomock.Any(), gomock.Any()).Return(&domain.Application{}, nil)
			},
			args: args{
				ctx:  context.Background(),
				appl: &domain.Application{},
			},
			expected:    nil,
			expectedErr: domain.ErrApplicationAlreadyExists,
		},
		{
			name: "error storing application",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByTgUserID(gomock.Any(), gomock.Any()).Return(nil, domain.ErrApplicationNotFound)
				f.storage.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil, domain.ErrInternal)
			},
			args: args{
				ctx:  context.Background(),
				appl: &domain.Application{},
			},
			expected:    nil,
			expectedErr: domain.ErrInternal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				storage: mock_application.NewMockApplicationStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := NewService(f.storage)
			got, err := s.CreateApplication(tt.args.ctx, tt.args.appl)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestApplicationService_GetAllApplications(t *testing.T) {
	type fields struct {
		storage *mock_application.MockApplicationStorage
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		expected    []*domain.Application
		expectedErr error
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetAll(gomock.Any()).Return([]*domain.Application{{}}, nil)
			},
			args: args{
				ctx: context.Background(),
			},
			expected:    []*domain.Application{{}},
			expectedErr: nil,
		},
		{
			name: "error not found",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetAll(gomock.Any()).Return(nil, domain.ErrApplicationNotFound)
			},
			args: args{
				ctx: context.Background(),
			},
			expected:    nil,
			expectedErr: domain.ErrApplicationNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				storage: mock_application.NewMockApplicationStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := NewService(f.storage)
			got, err := s.GetAllApplications(tt.args.ctx)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestApplicationService_GetUnresolvedApplications(t *testing.T) {
	type fields struct {
		storage *mock_application.MockApplicationStorage
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		expected    []*domain.Application
		expectedErr error
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetUnresolved(gomock.Any()).Return([]*domain.Application{{}}, nil)
			},
			args: args{
				ctx: context.Background(),
			},
			expected:    []*domain.Application{{}},
			expectedErr: nil,
		},
		{
			name: "error not found",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetUnresolved(gomock.Any()).Return(nil, domain.ErrApplicationNotFound)
			},
			args: args{
				ctx: context.Background(),
			},
			expected:    nil,
			expectedErr: domain.ErrApplicationNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				storage: mock_application.NewMockApplicationStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := NewService(f.storage)
			got, err := s.GetUnresolvedApplications(tt.args.ctx)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestApplicationService_ApproveApplication(t *testing.T) {
	type fields struct {
		storage *mock_application.MockApplicationStorage
	}
	type args struct {
		ctx           context.Context
		applicationID uint
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		expected    *domain.Application
		expectedErr error
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&domain.Application{}, nil)
				f.storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&domain.Application{}, nil)
			},
			args: args{
				ctx:           context.Background(),
				applicationID: 1,
			},
			expected:    &domain.Application{},
			expectedErr: nil,
		},
		{
			name: "application not found",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, domain.ErrApplicationNotFound)
			},
			args: args{
				ctx:           context.Background(),
				applicationID: 1,
			},
			expected:    nil,
			expectedErr: domain.ErrApplicationNotFound,
		},
		{
			name: "error updating",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&domain.Application{}, nil)
				f.storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, domain.ErrInternal)
			},
			args: args{
				ctx:           context.Background(),
				applicationID: 1,
			},
			expected:    nil,
			expectedErr: domain.ErrInternal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				storage: mock_application.NewMockApplicationStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := NewService(f.storage)
			got, err := s.ApproveApplication(tt.args.ctx, tt.args.applicationID)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestApplicationService_RejectApplication(t *testing.T) {
	type fields struct {
		storage *mock_application.MockApplicationStorage
	}
	type args struct {
		ctx           context.Context
		applicationID uint
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		expected    *domain.Application
		expectedErr error
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&domain.Application{}, nil)
				f.storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&domain.Application{}, nil)
			},
			args: args{
				ctx:           context.Background(),
				applicationID: 1,
			},
			expected:    &domain.Application{},
			expectedErr: nil,
		},
		{
			name: "application not found",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(nil, domain.ErrApplicationNotFound)
			},
			args: args{
				ctx:           context.Background(),
				applicationID: 1,
			},
			expected:    nil,
			expectedErr: domain.ErrApplicationNotFound,
		},
		{
			name: "error updating",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&domain.Application{}, nil)
				f.storage.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, domain.ErrInternal)
			},
			args: args{
				ctx:           context.Background(),
				applicationID: 1,
			},
			expected:    nil,
			expectedErr: domain.ErrInternal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				storage: mock_application.NewMockApplicationStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := NewService(f.storage)
			got, err := s.RejectApplication(tt.args.ctx, tt.args.applicationID)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestApplicationService_GetApprovedApplications(t *testing.T) {
	type fields struct {
		storage *mock_application.MockApplicationStorage
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		expected    []*domain.Application
		expectedErr error
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetApproved(gomock.Any()).Return([]*domain.Application{{}}, nil)
			},
			args: args{
				ctx: context.Background(),
			},
			expected:    []*domain.Application{{}},
			expectedErr: nil,
		},
		{
			name: "error not found",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetApproved(gomock.Any()).Return(nil, domain.ErrApplicationNotFound)
			},
			args: args{
				ctx: context.Background(),
			},
			expected:    nil,
			expectedErr: domain.ErrApplicationNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				storage: mock_application.NewMockApplicationStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := NewService(f.storage)
			got, err := s.GetApprovedApplications(tt.args.ctx)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestApplicationService_GetRejectedApplications(t *testing.T) {
	type fields struct {
		storage *mock_application.MockApplicationStorage
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		expected    []*domain.Application
		expectedErr error
	}{
		{
			name: "success",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetRejected(gomock.Any()).Return([]*domain.Application{{}}, nil)
			},
			args: args{
				ctx: context.Background(),
			},
			expected:    []*domain.Application{{}},
			expectedErr: nil,
		},
		{
			name: "error not found",
			prepare: func(f *fields) {
				f.storage.EXPECT().GetRejected(gomock.Any()).Return(nil, domain.ErrApplicationNotFound)
			},
			args: args{
				ctx: context.Background(),
			},
			expected:    nil,
			expectedErr: domain.ErrApplicationNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				storage: mock_application.NewMockApplicationStorage(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := NewService(f.storage)
			got, err := s.GetRejectedApplications(tt.args.ctx)
			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}
