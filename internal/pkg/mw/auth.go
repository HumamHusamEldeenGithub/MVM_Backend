package mw

import (
	"fmt"
	"mvm_backend/internal/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(auth service.IMVMAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < len(BEARER_SCHEMA) {
			fmt.Println("NO AUTH HEADER")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		_, err := auth.VerifyToken(tokenString, false)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
