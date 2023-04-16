package store

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) JoinRoom(roomId, userID string) error {
	filter := bson.M{"id": roomId}
	update := bson.M{"$addToSet": bson.M{"users": userID}}
	_, err := repository.roomsCollection.UpdateOne(repository.ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}
