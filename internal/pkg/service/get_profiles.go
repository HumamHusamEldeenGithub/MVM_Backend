package service

import (
	"mvm_backend/internal/pkg/model"
)

func (s *mvmService) GetProfiles(ids []string) ([]*model.User, error) {
	return s.store.GetProfiles(ids)
}
