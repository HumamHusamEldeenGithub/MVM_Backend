package store

import (
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repository *MVMRepository) GetFriends(userID string) (*model.Friends, error) {
	filter := bson.D{{Key: "id", Value: userID}}
	var friends model.Friends
	if err := repository.friendsCollection.FindOne(repository.ctx, filter).Decode(&friends); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return &friends, nil
}
