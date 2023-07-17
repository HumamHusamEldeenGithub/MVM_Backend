package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) UpsertAvatarSettings(userId string, settings map[int64]int64) error {

	filter := bson.M{"id": userId}
	update := bson.M{
		"$set": bson.M{
			"id":       userId,
			"settings": settings,
		},
	}

	// Perform the upsert operation
	opts := options.Update().SetUpsert(true)
	_, err := repository.avatarsCollection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}
