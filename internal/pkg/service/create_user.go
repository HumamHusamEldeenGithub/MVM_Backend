package service

import (
	"context"
	"mvm_backend/internal/pkg/model"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) CreateUser(ctx context.Context, user *model.User) error {
	hashedPassword, err := utils.HashMyPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	if err := s.store.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
