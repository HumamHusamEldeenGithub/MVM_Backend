package service

import (
	"mvm_backend/internal/pkg/model"
)

func (s *mvmService) GetUser(email string) (*model.User, error) {
	user, err := s.store.GetUser(email)
	if err != nil {
		return nil, err
	}
	return user, err
}
