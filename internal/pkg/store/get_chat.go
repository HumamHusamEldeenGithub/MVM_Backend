package store

import (
	"mvm_backend/internal/pkg/generated/mvmPb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repository *MVMRepository) GetChat(userId1, userId2 string) (*mvmPb.Chat, error) {
	idsToSearch := []string{userId1, userId2}
	filter := bson.M{
		"participants": bson.M{
			"$all": idsToSearch,
		},
	}

	var chat mvmPb.Chat
	if err := repository.chatsCollection.FindOne(repository.ctx, filter).Decode(&chat); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return &chat, nil
}
