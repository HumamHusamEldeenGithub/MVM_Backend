package service

import (
	"fmt"
	"log"
)

func (s *mvmService) CreateRoomInvitation(userId, roomId, recipientId string) error {
	room, err := s.store.GetRoom(roomId)
	if err != nil {
		return err
	}
	if room.OwnerId != userId {
		return fmt.Errorf("Not room owner")
	}

	if err := s.store.CreateRoomInvitation(roomId, recipientId); err != nil {
		return err
	}
	notification, err := s.CreateRoomInvitationNotification(userId, recipientId, roomId)
	if err != nil {
		log.Printf("error: %v", err)
	}

	_, err = s.CreateNotification(notification)
	if err != nil {
		log.Printf("error: %v", err)
	}

	NotificationsBroadcaster <- notification
	return nil
}
