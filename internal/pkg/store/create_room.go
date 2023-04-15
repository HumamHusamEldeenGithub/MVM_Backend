package store

import (
	"fmt"
	"mvm_backend/internal/pkg/generated/mvmPb"

	"go.mongodb.org/mongo-driver/mongo"
)

func (repository *MVMRepository) CreateRoom(room *mvmPb.Room) (*mvmPb.Room, error) {

	_, err := repository.roomsCollection.InsertOne(repository.ctx, room)
	if err != nil {
		// Check for duplicate key error
		writeException, ok := err.(mongo.WriteException)
		if ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 || writeError.Code == 11001 {
					return nil, fmt.Errorf("Room already exists ")
				} else {
					return nil, err
				}
			}
		} else {
			// Handle other types of errors
			return nil, err
		}
	}

	return room, nil
}
