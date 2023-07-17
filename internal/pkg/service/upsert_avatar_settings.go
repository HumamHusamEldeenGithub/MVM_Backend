package service

func (s *mvmService) UpsertAvatarSettings(userID string, settings map[int64]int64) error {
	return s.store.UpsertAvatarSettings(userID, settings)
}
