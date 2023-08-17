package service

import "mvm_backend/internal/pkg/generated/mvmPb"

func (s *mvmService) GetChat(userId1, userId2 string) (*mvmPb.Chat, error) {
	chat, err := s.store.GetChat(userId1, userId2)
	if err != nil {
		return nil, err
	}
	for i, j := 0, len(chat.Messages)-1; i < j; i, j = i+1, j-1 {
		chat.Messages[i], chat.Messages[j] = chat.Messages[j], chat.Messages[i]
	}
	return chat, err
}
