package memberships

import (
	"testing"

	"github.com/IndraNurfa/music-catalog/internal/configs"
	"github.com/IndraNurfa/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "sucess",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@example.com",
					Password: "123",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@example.com",
					Password: "$2a$10$MGx3xMwgXoerBX1ikHvoUuj6ByOyNvdUorUwmLBSyTxyaLD3nQxg.",
					Username: "test",
				}, nil)
			},
		},
		{
			name: "failed when get user",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@example.com",
					Password: "123",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when passwords are not match",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@example.com",
					Password: "123",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@example.com",
					Password: "wrong password",
					Username: "test",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJWT: "abcde",
					},
				},
				repository: mockRepo,
			}
			got, err := s.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
