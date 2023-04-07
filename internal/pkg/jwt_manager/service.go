package jwt_manager

import "time"

type authService struct {
	secret               string
	refreshSecret        string
	tokenDuration        time.Duration
	refreshTokenDuration time.Duration
}

func NewAuthService(secret, refreshSecret string, tokenDuration time.Duration, refreshTokenDuration time.Duration) *authService {
	return &authService{
		secret:               secret,
		refreshSecret:        refreshSecret,
		tokenDuration:        tokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}
}
