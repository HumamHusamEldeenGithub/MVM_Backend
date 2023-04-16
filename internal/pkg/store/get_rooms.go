package store

import (
	"mvm_backend/internal/pkg/generated/mvmPb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) GetRooms() ([]*mvmPb.Room, error) {
	cursor, err := repository.roomsCollection.Find(repository.ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}

	var rooms []*mvmPb.Room

	for cursor.Next(repository.ctx) {
		var room mvmPb.Room
		err := cursor.Decode(&room)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
