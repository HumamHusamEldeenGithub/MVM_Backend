package service

import (
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/utils"
)

func (s *mvmService) CreateNotification(notification *mvmPb.Notification) (*mvmPb.Notification, error) {
	notification.Id = utils.GenerateID()
	return s.store.CreateNotification(notification)
}
