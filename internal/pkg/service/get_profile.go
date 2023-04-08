package service

import (
	"mvm_backend/internal/pkg/model"
)

func (s *mvmService) GetProfile(id string) (*model.User, error) {
	user, err := s.store.GetProfile(id, false)
	if err != nil {
		return nil, err
	}
	return user, err
}
