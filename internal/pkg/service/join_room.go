package service

func (s *mvmService) JoinRoom(roomId, userId string) error {
	return s.store.JoinRoom(roomId, userId)
}
