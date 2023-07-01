package jwt_manager

import (
	"fmt"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	UserID string `json:"userid"`
	Role   string `json:"role"`
}

type JWTToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *authService) GenerateToken(user *model.User, generateRefreshToken bool) (*JWTToken, error) {
	claimsToken := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.tokenDuration).Unix(),
		},
		UserID: user.ID,
		Role:   "1",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsToken)
	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		fmt.Println("Error signing JWT token:", err)
		return nil, err
	}

	var signedRefreshToken string
	if generateRefreshToken {
		claimsRefreshToken := UserClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(s.refreshTokenDuration).Unix(), // Set token expiration to 24 hours from now
			},
			UserID: user.ID,
			Role:   "1",
		}

		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefreshToken)
		signedRefreshToken, err = refreshToken.SignedString([]byte(s.refreshSecret))
		if err != nil {
			fmt.Println("Error signing JWT refresh token:", err)
			return nil, err
		}
	}

	return &JWTToken{
		Token:        signedToken,
		RefreshToken: signedRefreshToken,
	}, nil
}

func (s *authService) VerifyToken(tokenString string, isRefreshToken bool) (string, error) {
	claims, err := s.GetUserClaims(tokenString, isRefreshToken)
	if err != nil {
		return "", err
	}

	return claims.UserID, nil
}

func (s *authService) GetUserClaims(tokenString string, isRefreshToken bool) (*UserClaims, error) {
	var secret string
	if isRefreshToken {
		secret = s.refreshSecret
	} else {
		secret = s.secret
	}

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("Error parsing JWT token:", err)
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			return nil, errors.ErrorsList[errors.ExpiredTokenError]
		}
		return nil, errors.NewErrorDesc(err.Error(), http.StatusUnauthorized)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.NewErrorDesc("invalid JWT token", http.StatusUnauthorized)
	}
	return claims, nil
}
