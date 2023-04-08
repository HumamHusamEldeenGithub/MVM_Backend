package service

func (s *mvmService) AddFriend(userID, friendID string) error {
	return s.store.AddFriend(userID, friendID)
}
