package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"strconv"
	"{{cookiecutter.project_name}}/store/sqlq"
)

type Context struct {
	Db      *sql.DB
	Pool    *pgxpool.Pool
	Queries *sqlq.Queries
	UserId  int64
	Logger  logrus.FieldLogger
}

func GetIdFromRequest(c *gin.Context, key string) int64 {
	// Get ID from URL
	idStr := c.Param(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

func ContextMiddleware(db *sql.DB, pool *pgxpool.Pool, logger logrus.FieldLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ctx", Context{
			Db:      db,
			Pool:    pool,
			Queries: sqlq.New(pool),
			UserId:  c.GetInt64("user_id"),
			Logger:  logger,
		})
		c.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.MustGet("ctx").(Context).UserId == 0 {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
