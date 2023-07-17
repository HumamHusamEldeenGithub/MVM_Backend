package service

func (s *mvmService) GetAvatarSettings(id string) (map[int32]int32, error) {
	return s.store.GetAvatarSettings(id)
}
