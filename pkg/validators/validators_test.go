package validators

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestApplicationService_IsValidPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			name: "valid phone",
			args: args{
				phone: "+79999999999",
			},
			expected: true,
		},
		{
			name: "empty string",
			args: args{
				phone: "",
			},
			expected: false,
		},
		{
			name: "no plus",
			args: args{
				phone: "89999999999",
			},
			expected: false,
		},
		{
			name: "wrong digits amount",
			args: args{
				phone: "7777777777",
			},
			expected: false,
		},
		{
			name: "random string",
			args: args{
				phone: "asdasd",
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			got := IsValidPhone(tt.args.phone)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestApplicationService_IsValidFullName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			name: "valid cyrillic name",
			args: args{
				name: "Иванов Иван Иванович",
			},
			expected: true,
		},
		{
			name: "valid latinic name",
			args: args{
				name: "Ivanov Ivan Ivanovich",
			},
			expected: true,
		},
		{
			name: "empty string",
			args: args{
				name: "",
			},
			expected: false,
		},
		{
			name: "random string",
			args: args{
				name: "asdasdasd",
			},
			expected: false,
		},
		{
			name: "non-full name",
			args: args{
				name: "Иван Иванов",
			},
			expected: false,
		},
		{
			name: "three spaces",
			args: args{
				name: "   ",
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			got := IsValidFullName(tt.args.name)
			assert.Equal(t, tt.expected, got)
		})
	}
}
