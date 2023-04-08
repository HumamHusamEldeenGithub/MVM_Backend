package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) DeleteFriend(userID, friendID string) error {
	filter := bson.M{"id": userID}
	update := bson.M{"$pull": bson.M{"friends": friendID}}
	result, err := repository.usersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	filter = bson.M{"id": friendID}
	update = bson.M{"$pull": bson.M{"friends": userID}}
	result, err = repository.usersCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println(result.MatchedCount)

	return nil
}
