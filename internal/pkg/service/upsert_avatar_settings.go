package service

import "mvm_backend/internal/pkg/generated/mvmPb"

func (s *mvmService) UpsertAvatarSettings(userID string, settings *mvmPb.AvatarSettings) error {
	return s.store.UpsertAvatarSettings(userID, settings)
}
