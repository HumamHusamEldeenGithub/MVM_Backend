package store

import (
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) GetProfile(id string, withPassword bool) (*model.User, error) {
	filter := bson.D{{Key: "id", Value: id}}

	var user model.User
	var avatarSettings model.AvatarSettings

	var opt *options.FindOneOptions
	if !withPassword {
		opt = options.FindOne().SetProjection(bson.M{"password": 0})
	}

	if err := repository.usersCollection.FindOne(repository.ctx, filter, opt).Decode(&user); err != nil {
		return nil, err
	}

	if err := repository.avatarsCollection.FindOne(repository.ctx, filter).Decode(&avatarSettings); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}

	user.AvatarSettings = avatarSettings.Settings

	return &user, nil
}
