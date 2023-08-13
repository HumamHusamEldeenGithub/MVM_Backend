package service

import "mvm_backend/internal/pkg/model"

func (s *mvmService) GetFriends(id string) (*model.Friends, error) {
	user, err := s.store.GetFriends(id)
	if err != nil {
		return nil, err
	}
	return user, err
}
