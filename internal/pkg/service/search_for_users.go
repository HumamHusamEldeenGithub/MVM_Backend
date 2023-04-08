package service

import "mvm_backend/internal/pkg/model"

func (s *mvmService) SearchForUsers(searchInput string) ([]*model.User, error) {
	users, err := s.store.SearchForUsers(searchInput)
	if err != nil {
		return nil, err
	}
	return users, nil
}
