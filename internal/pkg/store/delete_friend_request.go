package store

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) DeleteFriendRequest(userID, friendID string) error {
	filter := bson.M{"id": friendID}
	update := bson.M{"$pull": bson.M{"pending": userID}}
	_, err := repository.friendsCollection.UpdateOne(repository.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
