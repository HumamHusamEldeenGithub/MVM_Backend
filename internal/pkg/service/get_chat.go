package service

import "mvm_backend/internal/pkg/generated/mvmPb"

func (s *mvmService) GetChat(userId1, userId2 string) (*mvmPb.Chat, error) {
	chat, err := s.store.GetChat(userId1, userId2)
	if err != nil {
		return nil, err
	}
	return chat, err
}
