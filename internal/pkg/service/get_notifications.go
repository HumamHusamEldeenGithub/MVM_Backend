package service

import "mvm_backend/internal/pkg/generated/mvmPb"

func (s *mvmService) GetNotifications(id string) ([]*mvmPb.Notification, error) {
	notifications, err := s.store.GetNotifications(id)
	if err != nil {
		return nil, err
	}
	return notifications, err
}
