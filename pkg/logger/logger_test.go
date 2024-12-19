package logger

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestApplicationService_GetUnresolvedApplications(t *testing.T) {
	type fields struct{}
	type args struct {
		ctx         context.Context
		outputPaths []string
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		expectedErr error
	}{
		{
			name:    "success",
			prepare: func(f *fields) {},
			args: args{
				ctx:         context.Background(),
				outputPaths: []string{"stderr"},
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			_, err := NewZapLogger(tt.args.outputPaths)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
