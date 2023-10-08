package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"morgan.io/internal/platform/reqctx"
)

// AuthMiddleware is a middleware that checks if the request is authenticated
func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the request is authenticated
			claims, err := authenticate(r, secret)
			if err != nil {
				// If not authenticated, return a 401 Unauthorized response
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			userID, err := uuid.Parse(claims.Id)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Set the claims to the request context
			r = r.WithContext(reqctx.WithUserID(r.Context(), userID))

			// If authenticated, call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// authenticate is a function that checks if the request is authenticated
func authenticate(r *http.Request, secret string) (*Claims, error) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")

	// Check if the Authorization header is empty
	if authHeader == "" {
		return nil, errors.New("authorization header is empty")
	}

	// Get bearer key from the Authorization header
	key := strings.Split(authHeader, "Bearer ")[1]

	token := Token(key)

	return token.Validate(secret)
}
