package service

import "fmt"

func (s *mvmService) DeleteRoomInvitation(userId, roomId, recipientId string) error {
	room, err := s.store.GetRoom(roomId)
	if err != nil {
		return err
	}
	if room.OwnerId != userId {
		return fmt.Errorf("Not room owner")
	}
	return s.store.DeleteRoomInvitation(roomId, recipientId)
}
