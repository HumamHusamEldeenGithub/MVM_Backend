package mw

import (
	"context"
	"mvm_backend/internal/pkg/service"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//	func AuthorizeJWT(auth service.IMVMAuth) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			const BEARER_SCHEMA = "Bearer "
//			authHeader := c.GetHeader("Authorization")
//			if len(authHeader) < len(BEARER_SCHEMA) {
//				fmt.Println("NO AUTH HEADER")
//				c.AbortWithStatus(http.StatusUnauthorized)
//				return
//			}
//			tokenString := authHeader[len(BEARER_SCHEMA):]
//			_, err := auth.VerifyToken(tokenString, false)
//			if err != nil {
//				fmt.Println(err)
//				c.AbortWithStatus(http.StatusUnauthorized)
//			}
//		})
//	}
// func AuthorizeJWT(arg1 string, arg2 int) mux.MiddlewareFunc {
// 	return func(next http.Handler, auth service.IMVMAuth) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			// Get the JWT token from the Authorization header
// 			authHeader := r.Header.Get("Authorization")
// 			tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")

// 			userID, err := auth.VerifyToken(tokenString, false)
// 			if err != nil {
// 				http.Error(w, "Invalid token", http.StatusUnauthorized)
// 				return
// 			}

// 			// Attach the user ID to the request context
// 			ctx := context.WithValue(r.Context(), "user_id", userID)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		})
// 	}
// }

func MyMiddleware(auth service.IMVMAuth) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the JWT token from the Authorization header
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")

			userID, err := auth.VerifyToken(tokenString, false)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Attach the user ID to the request context
			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// // Middleware to authenticate requests with JWT tokens
// func AuthorizeJWT(next http.Handler, auth service.IMVMAuth) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Get the JWT token from the Authorization header
// 		authHeader := r.Header.Get("Authorization")
// 		tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")

// 		userID, err := auth.VerifyToken(tokenString, false)
// 		if err != nil {
// 			http.Error(w, "Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		// Attach the user ID to the request context
// 		ctx := context.WithValue(r.Context(), "user_id", userID)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
