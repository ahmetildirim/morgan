package auth

import (
	"net/http"
	"strings"
)

// AuthMiddleware is a middleware that checks if the request is authenticated
func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the request is authenticated
			if !isAuthenticated(r, secret) {
				// If not authenticated, return a 401 Unauthorized response
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// If authenticated, call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// isAuthenticated is a function that checks if the request is authenticated
func isAuthenticated(r *http.Request, secret string) bool {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")

	// Check if the Authorization header is empty
	if authHeader == "" {
		return false
	}

	// Get bearer token from the Authorization header
	token := strings.Split(authHeader, "Bearer ")[1]

	_, err := ValidateToken(token, secret)
	if err != nil {
		return false
	}

	return true
}
