package service

import (
	"mvm_backend/internal/pkg/model"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) CreateUser(user *model.User) (string, error) {
	hashedPassword, err := utils.HashMyPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	userID, err := s.store.CreateUser(user)
	if err != nil {
		return "", err
	}
	return userID, nil
}
