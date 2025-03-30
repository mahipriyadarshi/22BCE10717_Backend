package config

import (
	"context"
	"fmt"
	"net/http"
	"strings" // Add missing import
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("your_secure_secret_key")

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx := context.WithValue(r.Context(), "email", claims["email"].(string))
			next.ServeHTTP(w, r.WithContext(ctx)) // Fixed context handling
			return
		}
		
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
