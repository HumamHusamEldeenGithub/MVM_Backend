package store

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (repository *MVMRepository) DeleteNotification(userId, notificationId string) error {
	filter := bson.M{"userid": userId, "id": notificationId}
	_, err := repository.notificationsCollection.DeleteOne(repository.ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
