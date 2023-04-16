package service

import "mvm_backend/internal/pkg/generated/mvmPb"

func (s *mvmService) GetRoom(id string) (*mvmPb.Room, error) {
	return s.store.GetRoom(id)
}
