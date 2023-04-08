package service

import (
	"mvm_backend/internal/pkg/model"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) CreateUser(user *model.User) (string, error) {
	user.ID = utils.GenerateID()
	user.Friends = []string{}

	hashedPassword, err := utils.HashMyPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	_, err = s.store.CreateUser(user)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}
