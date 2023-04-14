package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) CreateFriendRequest(userID, friendID string) error {
	filter := bson.M{"id": friendID}
	update := bson.M{"$addToSet": bson.M{"pending": userID}}
	_, err := repository.friendsCollection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}
