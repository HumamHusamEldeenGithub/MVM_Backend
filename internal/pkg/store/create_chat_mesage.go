package store

import (
	"mvm_backend/internal/pkg/generated/mvmPb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) CreateChatMessage(id string, chatMsg *mvmPb.ChatMessage) error {
	filter := bson.M{"id": id}
	update := bson.M{"$addToSet": bson.M{"messages": chatMsg}}
	_, err := repository.chatsCollection.UpdateOne(repository.ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}
