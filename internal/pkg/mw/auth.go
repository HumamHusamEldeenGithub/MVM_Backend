package mw

import (
	"fmt"
	"mvm_backend/internal/pkg/jwt_manager"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < len(BEARER_SCHEMA) {
			fmt.Println("NO AUTH HEADER")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		_, err := jwt_manager.VerifyToken(os.Getenv("JWT_SECRET"), tokenString)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
