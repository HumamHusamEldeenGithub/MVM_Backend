package service

func (s *mvmService) GetFriends(id string) ([]string, error) {
	user, err := s.store.GetFriends(id)
	if err != nil {
		return nil, err
	}
	return user, err
}
