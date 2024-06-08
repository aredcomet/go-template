package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"{{.ProjectName}}/internal/hasher"
	"{{.ProjectName}}/internal/jwttoken"
	"{{.ProjectName}}/internal/middleware"
	"{{.ProjectName}}/store/sqlq"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context, secret string) {
	ctx := c.MustGet("ctx").(middleware.Context)
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		ctx.Logger.WithError(err).Error("Invalid request")
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	user, err := ctx.Queries.GetUserByUsername(c, req.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("User not found")
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := hasher.NewHasher().Compare(req.Password, user.Password); err != nil {
		ctx.Logger.WithError(err).Error("Password comparison failed")
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, expiresAt, issuedAt, err := jwttoken.GenerateTokenPair(user.ID, secret)
	if err != nil {
		ctx.Logger.WithError(err).Error("Token generation failed")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	err = ctx.Queries.SaveAuthTokenPair(c, sqlq.SaveAuthTokenPairParams{
		UserID:    user.ID,
		Token:     token.AccessToken,
		ExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: issuedAt, Valid: true},
		Token_2:   token.RefreshToken,
	})

	if err != nil {
		ctx.Logger.WithError(err).Error("Token save failed")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, token)
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh" validate:"required"`
}

func Refresh(c *gin.Context, secret string) {
	ctx := c.MustGet("ctx").(middleware.Context)
	var req RefreshRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	authToken, err := ctx.Queries.GetActiveRefreshToken(c, req.RefreshToken)
	if err != nil {
		ctx.Logger.WithError(err).Error("Refresh token is invalid")
		c.JSON(401, gin.H{"error": "Refresh token is invalid"})
		return
	}

	if authToken.ExpiresAt.Time.Before(time.Now()) { // Refresh token has expired
		c.JSON(400, gin.H{"error": "Refresh token has expired"})
		return
	}

	token, expiresAt, issuedAt, err := jwttoken.GenerateTokenPair(authToken.UserID, secret)
	if err != nil {
		ctx.Logger.WithError(err).Error("Token generation failed")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err = ctx.Queries.SaveAuthTokenPair(c, sqlq.SaveAuthTokenPairParams{
		UserID:    authToken.UserID,
		Token:     token.AccessToken,
		ExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: issuedAt, Valid: true},
		Token_2:   token.RefreshToken,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, token)
}

func Me(c *gin.Context) {
	ctx := c.MustGet("ctx").(middleware.Context)
	user, err := ctx.Queries.GetUserById(c, ctx.UserId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(200, user)
}
