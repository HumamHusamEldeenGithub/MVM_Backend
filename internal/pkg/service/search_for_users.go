package service

import "mvm_backend/internal/pkg/model"

func (s *mvmService) SearchForUsers(searchInput, userId string) ([]*model.User, error) {
	users, err := s.store.SearchForUsers(searchInput, userId)
	if err != nil {
		return nil, err
	}

	return users, nil
}
