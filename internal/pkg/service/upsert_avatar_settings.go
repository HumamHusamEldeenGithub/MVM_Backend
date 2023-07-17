package service

func (s *mvmService) UpsertAvatarSettings(userID string, settings map[int32]int32) error {
	return s.store.UpsertAvatarSettings(userID, settings)
}
