package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) DeleteFriendRequest(userID, friendID string) error {
	filter := bson.M{"id": friendID}
	update := bson.M{"$pull": bson.M{"pending": userID}}
	_, err := repository.friendsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
