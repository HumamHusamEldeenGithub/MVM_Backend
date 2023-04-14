package service

func (s *mvmService) DeleteFriendRequest(userID, friendID string) error {
	return s.store.DeleteFriendRequest(userID, friendID)
}
