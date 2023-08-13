package store

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) CreateFriendRequest(userID, friendID string) error {
	filter := bson.M{"id": friendID}
	update := bson.M{"$addToSet": bson.M{"pending": userID}}
	_, err := repository.friendsCollection.UpdateOne(repository.ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	filter = bson.M{"id": userID}
	update = bson.M{"$addToSet": bson.M{"sent": friendID}}
	_, err = repository.friendsCollection.UpdateOne(repository.ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}
