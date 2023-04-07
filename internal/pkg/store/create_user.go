package store

import (
	"context"
	"fmt"
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repository *MVMRepository) CreateUser(ctx context.Context, user *model.User) error {
	userDB := repository.mongoDBClient.Database("public").Collection("users")

	result, err := userDB.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	stringObjectID := result.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(stringObjectID)
	return nil
}
