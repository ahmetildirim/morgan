package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"morgan.io/internal/user"
)

const expireDuration = 72 * time.Hour

type Token string

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewToken(user *user.User, secretKey string) (*Token, error) {
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
		return nil, err
	}

	t := Token(tokenString)

	return &t, nil
}

func (t Token) String() string {
	return string(t)
}

func (t Token) Validate(secretKey string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(t.String(), &Claims{}, func(token *jwt.Token) (interface{}, error) { return []byte(secretKey), nil })
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
