package jwt_manager

import (
	"fmt"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	UserID string `json:"userid"`
	Role   string `json:"role"`
}

func (s *authService) GenerateToken(user *model.User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * 24).Unix(), // Set token expiration to 24 hours from now
		},
		UserID: user.ID.Hex(),
		Role:   "1",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Replace "secret" with your actual JWT secret key
	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		fmt.Println("Error signing JWT token:", err)
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(secret, tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Replace "secret" with your actual JWT secret key
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("Error parsing JWT token:", err)
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			fmt.Println("JWT token expired")
			return nil, errors.Errorf(errors.ExpiredTokenError)
		}
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		fmt.Println("Invalid JWT token")
		return nil, fmt.Errorf("invalid JWT token")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		fmt.Println("JWT token expired")
		return nil, fmt.Errorf("JWT token expired")
	}

	return claims, nil
}
