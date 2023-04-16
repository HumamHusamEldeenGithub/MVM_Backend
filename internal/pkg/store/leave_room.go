package store

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) LeaveRoom(roomId, userID string) error {
	filter := bson.M{"id": roomId}
	update := bson.M{"$pull": bson.M{"users": userID}}
	_, err := repository.roomsCollection.UpdateOne(repository.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
