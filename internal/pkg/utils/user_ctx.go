package utils

import (
	"context"

	"github.com/gin-gonic/gin"
)

type userIDKey struct{}

// GetUserID gets user id from context
func GetJWTToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	return authHeader[len("Bearer"):]
}

// UserIDToContext ...
func UserIDToContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}
