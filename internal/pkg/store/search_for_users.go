package store

import (
	"fmt"
	"mvm_backend/internal/pkg/model"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) SearchForUsers(searchInput string, userId string) ([]*model.User, error) {
	// Convert searchInput to lowercase
	lowercaseSearchInput := strings.ToLower(searchInput)
	filter := bson.M{
		"username": bson.M{"$regex": fmt.Sprintf(`.*%s.*`, lowercaseSearchInput), "$options": "i"},
		"id":       bson.M{"$ne": userId},
	}

	cur, err := repository.usersCollection.Find(repository.ctx, filter, options.Find().SetProjection(bson.M{"password": 0}))
	if err != nil {
		return nil, err
	}
	var users []*model.User
	for cur.Next(repository.ctx) {
		var user model.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
