package service

func (s *mvmService) CreateFriendRequest(userID, friendID string) error {
	return s.store.CreateFriendRequest(userID, friendID)
}
