package jwttoken

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"net/http"
)

// ErrWrongAlg is an error returned when the signing method is invalid.
var (
	ErrWrongAlg = errors.New("invalid signing method")
)

type ErrResponse struct {
	Error string `json:"error"`
}

func verifyToken(tokenString string, secret string, claims jwt.Claims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrWrongAlg
		}
		return []byte(secret), nil
	})
	return token, err
}

func Middleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := request.BearerExtractor{}.ExtractToken(c.Request)
		if err != nil {
			c.Next()
			return
		}

		token, err := verifyToken(tokenString, secret, &CustomClaims{})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrResponse{"access denied: token expired"})
			return
		}

		claims, ok := token.Claims.(*CustomClaims)
		if !ok || claims.TokenType != "access" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrResponse{"Unable to extract claims or token type is wrong"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
