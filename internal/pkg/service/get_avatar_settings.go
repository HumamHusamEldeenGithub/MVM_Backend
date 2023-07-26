package service

import "mvm_backend/internal/pkg/generated/mvmPb"

func (s *mvmService) GetAvatarSettings(id string) (*mvmPb.AvatarSettings, error) {
	return s.store.GetAvatarSettings(id)
}
