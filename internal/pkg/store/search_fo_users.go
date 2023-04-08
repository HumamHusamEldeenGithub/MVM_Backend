package store

import (
	"fmt"
	"mvm_backend/internal/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) SearchForUsers(searchInput string) ([]*model.User, error) {
	filter := bson.M{"username": bson.M{"$regex": fmt.Sprintf(`.*%s.*`, searchInput)}}

	cur, err := repository.usersCollection.Find(repository.ctx, filter)
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
