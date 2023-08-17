package service

import (
	"mvm_backend/internal/pkg/generated/mvmPb"
)

func (s *mvmService) CreateChatMessage(chatId string, message *mvmPb.ChatMessage) error {
	return s.store.CreateChatMessage(chatId, message)
}
