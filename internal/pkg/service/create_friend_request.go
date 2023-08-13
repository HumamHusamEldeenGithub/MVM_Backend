package service

import "log"

func (s *mvmService) CreateFriendRequest(userID, friendID string) error {
	if err := s.store.CreateFriendRequest(userID, friendID); err != nil {
		return nil
	}

	notification, err := s.CreateFriendRequestNotification(userID, friendID)
	if err != nil {
		log.Printf("error: %v", err)
	}

	_, err = s.CreateNotification(notification)
	if err != nil {
		log.Printf("error: %v", err)
	}

	NotificationsBroadcaster <- *notification

	return nil
}
