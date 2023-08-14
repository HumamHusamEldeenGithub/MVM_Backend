package service

import "log"

func (s *mvmService) AddFriend(userID, friendID string) error {
	if err := s.store.AddFriend(userID, friendID); err != nil {
		return err
	}

	notification, err := s.CreateAcceptFriendRequestNotification(userID, friendID)
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
