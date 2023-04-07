package service

import (
	"context"
	"mvm_backend/internal/pkg/model"
)

type IMVMStore interface {
	GetUser(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
}

type IMVMAuth interface {
	GenerateToken(user *model.User) (string, error)
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
