package service

import "fmt"

func (s *mvmService) CreateRoomInvitation(userId, roomId, recipientId string) error {
	room, err := s.store.GetRoom(roomId)
	if err != nil {
		return err
	}
	if room.OwnerId != userId {
		return fmt.Errorf("Not room owner")
	}
	return s.store.CreateRoomInvitation(roomId, recipientId)
}
