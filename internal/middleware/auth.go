package middleware

import (
	"context"
	"fmt"
	"jagajkn/internal/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			claims, err := utils.ValidateToken(tokenString, secretKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "userID", claims.UserID)
			ctx = context.WithValue(ctx, "nik", claims.NIK)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIDFromContext(r *http.Request) (string, error) {
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		return "", fmt.Errorf("userID not found in context")
	}
	return userID, nil
}

func GetNIKFromContext(r *http.Request) (string, error) {
	nik, ok := r.Context().Value("nik").(string)
	if !ok {
		return "", fmt.Errorf("NIK not found in context")
	}
	return nik, nil
}