package service

import (
	"mvm_backend/internal/pkg/jwt_manager"
)

func (s *mvmService) LoginByRefreshToken(refreshToken string) (string, *jwt_manager.JWTToken, error) {
	userID, err := s.auth.VerifyToken(refreshToken, true)
	if err != nil {
		return "", nil, err
	}
	user, err := s.store.GetProfile(userID, false)
	if err != nil {
		return "", nil, err
	}
	tokens, err := s.auth.GenerateToken(user, true)
	if err != nil {
		return "", nil, err
	}
	return userID, tokens, nil
}
