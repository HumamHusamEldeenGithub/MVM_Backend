package service

func (s *mvmService) DeleteNotification(userID, notificationID string) error {
	return s.store.DeleteNotification(userID, notificationID)
}
