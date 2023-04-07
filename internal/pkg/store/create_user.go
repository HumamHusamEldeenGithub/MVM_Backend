package store

import (
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repository *MVMRepository) CreateUser(user *model.User) (string, error) {
	userDB := repository.mongoDBClient.Database("public").Collection("users")

	result, err := userDB.InsertOne(repository.ctx, user)
	if err != nil {
		return "", err
	}
	stringObjectID := result.InsertedID.(primitive.ObjectID).Hex()
	return stringObjectID, nil
}
