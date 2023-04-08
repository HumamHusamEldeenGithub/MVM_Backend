package service

func (s *mvmService) DeleteFriend(userID, friendID string) error {
	return s.store.DeleteFriend(userID, friendID)
}
