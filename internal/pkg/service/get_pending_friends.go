package service

func (s *mvmService) GetPendingFriends(id string) ([]string, error) {
	user, err := s.store.GetPendingFriends(id)
	if err != nil {
		return nil, err
	}
	return user, err
}
