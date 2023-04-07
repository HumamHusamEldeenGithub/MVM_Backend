package service

import (
	"context"
	"mvm_backend/internal/pkg/model"
)

func (s *mvmService) GetUser(ctx context.Context, email string) (*model.User, error) {
	user, err := s.store.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, err
}
