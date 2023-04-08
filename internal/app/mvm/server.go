package mvm

import (
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/model"
	"mvm_backend/internal/pkg/payloads"
)

type MVMServiceServer struct {
	service IMVMService
}

type IMVMService interface {
	LoginUser(req *payloads.LoginUserRequest) (*jwt_manager.JWTToken, error)
	LoginByRefreshToken(refreshToken string) (*jwt_manager.JWTToken, error)
	CreateUser(user *model.User) (string, error)

	GetUserByUsername(username string) (*model.User, error)
	GetProfile(id string) (*model.User, error)
	SearchForUsers(searchInput string) ([]*model.User, error)

	AddFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error
}

func NewIMVMServiceServer(service IMVMService) *MVMServiceServer {
	return &MVMServiceServer{
		service: service,
	}
}
