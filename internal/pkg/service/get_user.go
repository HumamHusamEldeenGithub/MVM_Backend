package service

import (
	"mvm_backend/internal/pkg/model"
)

func (s *mvmService) GetUserByUsername(username string) (*model.User, error) {
	user, err := s.store.GetUserByUsername(username, false)
	if err != nil {
		return nil, err
	}
	return user, err
}
