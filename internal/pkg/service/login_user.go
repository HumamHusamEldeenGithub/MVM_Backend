package service

import (
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/payloads"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) LoginUser(req *payloads.LoginUserRequest) (*jwt_manager.JWTToken, error) {
	user, err := s.store.GetUserByUsername(req.Username, true)
	if err != nil {
		return nil, err
	}

	if !utils.ComparePasswords(user.Password, req.Password) {
		return nil, errors.Errorf(errors.InvalidPasswordError)
	}
	tokens, err := s.auth.GenerateToken(user, true)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
