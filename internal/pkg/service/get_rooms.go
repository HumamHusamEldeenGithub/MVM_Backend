package service

import "mvm_backend/internal/pkg/generated/mvmPb"

func (s *mvmService) GetRooms() ([]*mvmPb.Room, error) {
	return s.store.GetRooms()
}
