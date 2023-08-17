package service

import (
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) CreateChat(userId1, userId2 string) (string, error) {
	chat := &mvmPb.Chat{
		Id:           utils.GenerateID(),
		Participants: []string{userId1, userId2},
		Messages:     []*mvmPb.ChatMessage{},
	}
	return chat.Id, s.store.CreateChat(chat)
}
