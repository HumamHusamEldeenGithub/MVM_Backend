package mvm

import (
	"context"
	v1 "mvm_backend/internal/pkg/generated/mvm-api/v1"
	"mvm_backend/internal/pkg/model"
)

type MVMServiceServer struct {
	service IMVMService
	v1.UnimplementedMVMServiceServer
}

type IMVMService interface {
	LoginUser(ctx context.Context, email, password string) (string, error)
	GetUser(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
}

func NewIMVMServiceServer(service IMVMService) *MVMServiceServer {
	return &MVMServiceServer{
		service: service,
	}
}
