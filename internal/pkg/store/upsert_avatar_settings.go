package store

import (
	"context"
	"mvm_backend/internal/pkg/generated/mvmPb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) UpsertAvatarSettings(userId string, settings *mvmPb.AvatarSettings) error {
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
