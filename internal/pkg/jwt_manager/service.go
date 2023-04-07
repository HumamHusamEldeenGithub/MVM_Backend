package jwt_manager

import "time"

type authService struct {
	secret        string
	tokenDuration time.Duration
}

func NewAuthService(secret string, tokenDuration time.Duration) *authService {
	return &authService{
		secret:        secret,
		tokenDuration: tokenDuration,
	}
}
