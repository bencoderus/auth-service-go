package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_TOKEN_EXPIRES_IN_HOUR = 3

type JwtToken struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

func SignToken(username uint) (JwtToken, error) {
	secretKey := os.Getenv("JWT_SECRET")

	if secretKey == "" {
		return JwtToken{}, errors.New("unable to retrieve JWT secret")
	}
	expiresAt := time.Now().Add(time.Hour * time.Duration(JWT_TOKEN_EXPIRES_IN_HOUR)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      expiresAt,
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return JwtToken{}, err
	}

	return JwtToken{
		Token:     tokenString,
		ExpiresAt: expiresAt,
	}, nil
}

func ExtractJwtTokenClaim(token jwt.Token) jwt.MapClaims {
	claims := token.Claims.(jwt.MapClaims)

	return claims
}

func ParseToken(tokenString string) (jwt.Token, error) {
	secretKey := os.Getenv("JWT_SECRET")

	if secretKey == "" {
		return jwt.Token{}, errors.New("unable to retrieve JWT secret")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	return *token, err
}
