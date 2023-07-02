package service

import "mvm_backend/internal/pkg/generated/mvmPb"

func (s *mvmService) GetUserRooms(userId string) ([]*mvmPb.Room, error) {
	return s.store.GetUserRooms(userId)
}
