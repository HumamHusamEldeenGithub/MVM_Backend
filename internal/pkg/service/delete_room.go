package service

func (s *mvmService) DeleteRoom(userId, roomId string) error {
	return s.store.DeleteRoom(roomId)
}
