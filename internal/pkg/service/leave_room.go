package service

func (s *mvmService) LeaveRoom(roomId, userId string) error {
	return s.store.LeaveRoom(roomId, userId)
}
