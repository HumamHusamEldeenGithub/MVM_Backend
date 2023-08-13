package store

import (
	"mvm_backend/internal/pkg/generated/mvmPb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repository *MVMRepository) GetNotifications(userId string) ([]*mvmPb.Notification, error) {
	filter := bson.D{{Key: "userid", Value: userId}}

	cursor, err := repository.notificationsCollection.Find(repository.ctx, filter, options.Find())
	if err != nil {
		return nil, err
	}

	var notifications []*mvmPb.Notification

	for cursor.Next(repository.ctx) {
		var notification mvmPb.Notification
		err := cursor.Decode(&notification)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &notification)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}
