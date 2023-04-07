package store

import (
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) GetUser(email string) (*model.User, error) {
	userDB := repository.mongoDBClient.Database("public").Collection("users")

	filter := bson.D{{Key: "email", Value: email}}

	// execute the query to find the matching user
	var user model.User

	if err := userDB.FindOne(repository.ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
