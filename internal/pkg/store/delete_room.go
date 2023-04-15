package store

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) DeleteRoom(roomId string) error {
	filter := bson.M{"id": roomId}
	_, err := repository.roomsCollection.DeleteOne(repository.ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
