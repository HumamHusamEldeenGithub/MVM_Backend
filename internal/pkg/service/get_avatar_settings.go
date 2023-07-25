package service

func (s *mvmService) GetAvatarSettings(id string) (map[int32]string, error) {
	return s.store.GetAvatarSettings(id)
}
