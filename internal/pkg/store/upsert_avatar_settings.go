package store

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) UpsertAvatarSettings(userId string, settings map[int32]string) error {
	filter := bson.M{"id": userId}
	update := bson.M{
		"$set": bson.M{
			"id":       userId,
			"settings": settings,
		},
	}

	opts := options.Update().SetUpsert(true)
	_, err := repository.avatarsCollection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}
