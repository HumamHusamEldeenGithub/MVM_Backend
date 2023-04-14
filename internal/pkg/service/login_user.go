package service

import (
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) LoginUser(username, password string) (*jwt_manager.JWTToken, error) {
	user, err := s.store.GetUserByUsername(username, true)
	if err != nil {
		return nil, err
	}

	if !utils.ComparePasswords(user.Password, password) {
		return nil, errors.Errorf(errors.InvalidPasswordError)
	}
	tokens, err := s.auth.GenerateToken(user, true)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
