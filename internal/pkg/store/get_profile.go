package store

import (
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) GetProfile(id string, withPassword bool) (*model.User, error) {
	userDB := repository.mongoDBClient.Database("public").Collection("users")

	filter := bson.D{{Key: "id", Value: id}}

	var user model.User

	var opt *options.FindOneOptions
	if !withPassword {
		opt = options.FindOne().SetProjection(bson.M{"password": 0})
	}

	if err := userDB.FindOne(repository.ctx, filter, opt).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
