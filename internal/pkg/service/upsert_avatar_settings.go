package service

func (s *mvmService) UpsertAvatarSettings(userID string, settings map[string]string) error {
	return s.store.UpsertAvatarSettings(userID, settings)
}
