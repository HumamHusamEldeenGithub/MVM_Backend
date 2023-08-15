package service

func (s *mvmService) DeleteNotification(userID, notificationId string) error {
	return s.store.DeleteNotification(userID, notificationId)
}
