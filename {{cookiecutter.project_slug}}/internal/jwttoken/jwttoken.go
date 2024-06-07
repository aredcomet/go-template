package jwttoken

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type CustomClaims struct {
	TokenType string `json:"token_type"`
	UserID    int64  `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenPair struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

func generateAccessTokenJWT(userID int64, secret string, expiresAt, issuedAt time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := &CustomClaims{
		TokenType: "access",
		UserID:    userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ID:        uuid.New().String(),
		},
	}
	token.Claims = claims

	// Generate encoded token and send it as a response.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

// generateToken returns a new random token of given length.
func generateToken(tokenLength int) (string, error) {
	b := make([]byte, tokenLength)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("could not generate token: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func GenerateTokenPair(userID int64, secret string) (TokenPair, time.Time, time.Time, error) {
	var token TokenPair
	expiresAt, issuedAt := time.Now().Add(time.Minute*30), time.Now()
	accessToken, err := generateAccessTokenJWT(userID, secret, expiresAt, issuedAt)
	if err != nil {
		return token, expiresAt, issuedAt, err
	}
	refreshToken, err := generateToken(32)
	if err != nil {
		return token, expiresAt, issuedAt, err
	}
	token = TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}
	return token, expiresAt, issuedAt, nil
}
