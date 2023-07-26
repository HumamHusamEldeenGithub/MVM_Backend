package store

import (
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repository *MVMRepository) GetAvatarSettings(userID string) (*mvmPb.AvatarSettings, error) {
	filter := bson.D{{Key: "id", Value: userID}}
	var avatarSettings model.AvatarSettings
	if err := repository.avatarsCollection.FindOne(repository.ctx, filter).Decode(&avatarSettings); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return avatarSettings.Settings, nil
}
