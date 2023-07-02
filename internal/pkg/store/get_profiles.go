package store

import (
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) GetProfiles(ids []string) ([]*model.User, error) {
	filter := bson.M{
		"id": bson.M{
			"$in": ids,
		},
	}

	cursor, err := repository.usersCollection.Find(repository.ctx,
		filter, options.Find().SetProjection(bson.M{"password": 0}))
	if err != nil {
		return nil, err
	}

	var users []*model.User

	for cursor.Next(repository.ctx) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
