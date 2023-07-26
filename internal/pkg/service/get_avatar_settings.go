package service

func (s *mvmService) GetAvatarSettings(id string) (map[string]string, error) {
	return s.store.GetAvatarSettings(id)
}
