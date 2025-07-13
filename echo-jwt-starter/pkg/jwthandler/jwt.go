package jwthandler

import (
	"echo-jwt-starter/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
)

type CustomClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type Payload struct {
	ID              string
	Role            string
	Subject         TokenType
	ExpirationHours int
}

// GenerateToken generates a new JWT token
func GenerateToken(p Payload) (string, error) {
	now := time.Now().UTC()

	claims := CustomClaims{
		ID:   p.ID,
		Role: p.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Envs.App.Name,
			Subject:   string(p.Subject),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(p.ExpirationHours) * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	signedToken, err := token.SignedString([]byte(config.Envs.Guard.JwtSecret))
	if err != nil {
		log.Error().Err(err).Msg("jwthandler::GenerateToken - signing failed")
		return "", err
	}

	return signedToken, nil
}

// ParseToken parses and validates JWT token string
func ParseToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Envs.Guard.JwtSecret), nil
	})
	if err != nil {
		log.Error().Err(err).Msg("jwthandler::ParseToken - parse failed")
		return nil, err
	}
	if !token.Valid {
		log.Warn().Msg("jwthandler::ParseToken - token is invalid")
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
