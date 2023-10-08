package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"morgan.io/internal/user"
)

const expireDuration = 72 * time.Hour

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func newAccessToken(user *user.User, secretKey string) (string, error) {
	claims := Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "morgan.io",
			Id:        user.ID.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}

func validateToken(tokenString string, secretKey string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) { return []byte(secretKey), nil })
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
