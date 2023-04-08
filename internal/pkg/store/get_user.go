package store

import (
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) GetUserByUsername(username string) (*model.User, error) {
	userDB := repository.mongoDBClient.Database("public").Collection("users")

	filter := bson.D{{Key: "username", Value: username}}

	var user model.User

	if err := userDB.FindOne(repository.ctx, filter, options.FindOne().SetProjection(bson.M{"password": 0})).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
