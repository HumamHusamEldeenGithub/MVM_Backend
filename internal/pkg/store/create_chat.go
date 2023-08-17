package store

import (
	"mvm_backend/internal/pkg/generated/mvmPb"
)

func (repository *MVMRepository) CreateChat(chat *mvmPb.Chat) error {
	_, err := repository.chatsCollection.InsertOne(repository.ctx, chat)
	if err != nil {
		return err

	}
	return nil
}
