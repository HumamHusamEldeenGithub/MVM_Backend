package store

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) DeleteFriend(userID, friendID string) error {
	filter := bson.M{"id": userID}
	update := bson.M{"$pull": bson.M{"friends": friendID}}
	_, err := repository.friendsCollection.UpdateOne(repository.ctx, filter, update)
	if err != nil {
		return err
	}

	filter = bson.M{"id": friendID}
	update = bson.M{"$pull": bson.M{"friends": userID}}
	_, err = repository.friendsCollection.UpdateOne(repository.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
