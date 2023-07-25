package service

func (s *mvmService) UpsertAvatarSettings(userID string, settings map[int32]string) error {
	return s.store.UpsertAvatarSettings(userID, settings)
}
