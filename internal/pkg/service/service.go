package service

import (
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/model"
)

type IMVMStore interface {
	GetUser(email string) (*model.User, error)
	CreateUser(user *model.User) (string, error)
}

type IMVMAuth interface {
	GenerateToken(user *model.User, refreshToken bool) (*jwt_manager.JWTToken, error)
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
