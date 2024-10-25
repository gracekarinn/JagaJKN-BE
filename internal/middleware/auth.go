package middleware

import (
	"context"
	"fmt"
	"jagajkn/internal/utils"
	"net/http"
	"strings"
)

type contextKey string

const (
    NIKKey contextKey = "nik"
)

func AuthMiddleware(secretKey string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            fmt.Println("Processing request in AuthMiddleware") 

            authHeader := r.Header.Get("Authorization")
            fmt.Printf("Auth Header: %s\n", authHeader) 

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
                fmt.Printf("Token validation error: %v\n", err) 
                http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
                return
            }

            fmt.Printf("Token claims NIK: %s\n", claims.NIK)

            ctx := r.Context()
            ctx = context.WithValue(ctx, NIKKey, claims.NIK)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func GetNIKFromContext(r *http.Request) (string, error) {
    nik, ok := r.Context().Value(NIKKey).(string)
    if !ok {
        return "", fmt.Errorf("NIK not found in context")
    }
    return nik, nil
}

