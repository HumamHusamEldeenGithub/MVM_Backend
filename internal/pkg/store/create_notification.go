package store

import (
	"mvm_backend/internal/pkg/generated/mvmPb"
)

func (repository *MVMRepository) CreateNotification(notification *mvmPb.Notification) (*mvmPb.Notification, error) {
	_, err := repository.notificationsCollection.InsertOne(repository.ctx, notification)
	if err != nil {
		return nil, err

	}
	return notification, nil
}
