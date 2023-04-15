package service

import (
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) CreateRoom(room *mvmPb.Room) (*mvmPb.Room, error) {
	room.Id = utils.GenerateID()
	return s.store.CreateRoom(room)
}
