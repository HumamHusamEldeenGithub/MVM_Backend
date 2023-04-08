package service

import (
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/model"
)

type IMVMStore interface {
	CreateUser(user *model.User) (string, error)

	GetProfile(id string, withPassword bool) (*model.User, error)
	GetUserByUsername(username string, withPassword bool) (*model.User, error)
	SearchForUsers(searchInput string) ([]*model.User, error)

	AddFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error
}

type IMVMAuth interface {
	GenerateToken(user *model.User, refreshToken bool) (*jwt_manager.JWTToken, error)
	VerifyToken(token string, isRefreshToken bool) (string, error)
}

type mvmService struct {
	store IMVMStore
	auth  IMVMAuth
}

func NewMVMService(store IMVMStore, auth IMVMAuth) *mvmService {
	return &mvmService{
		store: store,
		auth:  auth,
	}
}
