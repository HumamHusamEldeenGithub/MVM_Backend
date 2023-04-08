package store

import (
	"fmt"
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repository *MVMRepository) CreateUser(user *model.User) (string, error) {
	userDB := repository.mongoDBClient.Database("public").Collection("users")

	result, err := userDB.InsertOne(repository.ctx, user)
	if err != nil {
		// Check for duplicate key error
		writeException, ok := err.(mongo.WriteException)
		if ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 || writeError.Code == 11001 {
					return "", fmt.Errorf("Username or Email already exists ")
				} else {
					return "", err
				}
			}
		} else {
			// Handle other types of errors
			return "", err
		}
	}
	stringObjectID := result.InsertedID.(primitive.ObjectID).Hex()
	return stringObjectID, nil
}
