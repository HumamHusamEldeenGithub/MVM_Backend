package service

import (
	"mvm_backend/internal/pkg/model"
)

func (s *mvmService) GetProfile(id string) (*model.User, error) {
	return s.store.GetProfile(id, false)
}
