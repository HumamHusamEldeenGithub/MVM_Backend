package service

import "mvm_backend/internal/pkg/generated/mvmPb"

func (s *mvmService) GetRooms(searchQuery string) ([]*mvmPb.Room, error) {
	return s.store.GetRooms(searchQuery)
}
