package service

func (s *mvmService) GetAvatarSettings(id string) (map[int64]int64, error) {
	return s.store.GetAvatarSettings(id)
}
