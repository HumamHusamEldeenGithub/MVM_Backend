package store

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) DeleteNotifications(userId string) error {
	filter := bson.M{"userid": userId}
	_, err := repository.notificationsCollection.DeleteMany(repository.ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
