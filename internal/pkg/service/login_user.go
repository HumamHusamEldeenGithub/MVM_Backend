package service

import (
	"context"
	"fmt"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) LoginUser(ctx context.Context, email, password string) (string, error) {
	user, err := s.store.GetUser(ctx, email)
	if err != nil {
		return "", err
	}

	if !utils.ComparePasswords(user.Password, password) {
		return "", fmt.Errorf("incorrect password")
	}
	token, err := s.auth.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
