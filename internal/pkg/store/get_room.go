package store

import (
	"mvm_backend/internal/pkg/generated/mvmPb"

	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) GetRoom(id string) (*mvmPb.Room, error) {
	filter := bson.D{{Key: "id", Value: id}}
	var room mvmPb.Room
	if err := repository.roomsCollection.FindOne(repository.ctx, filter).Decode(&room); err != nil {
		return nil, err
	}
	return &room, nil
}
