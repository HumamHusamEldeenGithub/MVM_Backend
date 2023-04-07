package mvm

import (
	v1 "mvm_backend/internal/pkg/generated/mvm-api/v1"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/model"
	"mvm_backend/internal/pkg/payloads"
)

type MVMServiceServer struct {
	service IMVMService
	v1.UnimplementedMVMServiceServer
}

type IMVMService interface {
	LoginUser(req *payloads.LoginUserRequest) (*jwt_manager.JWTToken, error)
	GetUser(email string) (*model.User, error)
	CreateUser(user *model.User) (string, error)
}

func NewIMVMServiceServer(service IMVMService) *MVMServiceServer {
	return &MVMServiceServer{
		service: service,
	}
}
