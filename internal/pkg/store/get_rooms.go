package store

import (
	"fmt"
	"mvm_backend/internal/pkg/generated/mvmPb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) GetRooms(searchQuery string) ([]*mvmPb.Room, error) {
	var filter interface{} = bson.D{}
	if len(searchQuery) != 0 {
		filter = bson.M{"title": bson.M{"$regex": fmt.Sprintf(`(?i)^%s`, searchQuery)}}
	}
	cursor, err := repository.roomsCollection.Find(repository.ctx, filter, options.Find())
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
