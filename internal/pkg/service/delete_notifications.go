package service

func (s *mvmService) DeleteNotifications(userID string) error {
	return s.store.DeleteNotifications(userID)
}
