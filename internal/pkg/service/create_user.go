package service

import (
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/model"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) CreateUser(user *model.User) (string, *jwt_manager.JWTToken, error) {
	user.ID = utils.GenerateID()
	user.Friends = []string{}

	hashedPassword, err := utils.HashMyPassword(user.Password)
	if err != nil {
		return "", nil, err
	}
	user.Password = hashedPassword

	_, err = s.store.CreateUser(user)
	if err != nil {
		return "", nil, err
	}

	tokens, err := s.auth.GenerateToken(user, true)
	if err != nil {
		return "", nil, err
	}
	return user.ID, tokens, nil
}
