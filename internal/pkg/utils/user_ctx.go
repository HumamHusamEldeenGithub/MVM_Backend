package utils

import (
	"context"
)

type userIDKey struct{}

// GetUserID gets user id from context
func GetUserID(ctx context.Context) string {
	if v, ok := ctx.Value(userIDKey{}).(string); ok {
		return v
	}
	return ""
}

// UserIDToContext ...
func UserIDToContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}
