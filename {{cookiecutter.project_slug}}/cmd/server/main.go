package main

import (
	"database/sql"
	"log"
	"{{cookiecutter.project_name}}/config"
	"{{cookiecutter.project_name}}/internal/jwttoken"
	"{{cookiecutter.project_name}}/internal/middleware"
	"{{cookiecutter.project_name}}/pkg/auth"
	"{{cookiecutter.project_name}}/store/database"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.NewConfig()
	db := database.Connect(cfg.PostgresUri, cfg.SqlLogLevel, cfg.Logger)
	pool := database.ConnectPool(cfg.PostgresUri, cfg.SqlLogLevel, cfg.Logger)

	r := gin.Default()

	definePublicRoutes(r, cfg, db, pool)
	definePrivateRoutes(r, cfg, db, pool)

	log.Fatal(r.Run(":8080"))
}

func definePublicRoutes(r *gin.Engine, cfg *config.Config, db *sql.DB, pgxPool *pgxpool.Pool) {
	public := r.Group("/api/v1")
	public.Use(middleware.ContextMiddleware(db, pgxPool, cfg.Logger))
	public.POST("/token", loginHandler(cfg.JwtSecret))
}

func definePrivateRoutes(r *gin.Engine, cfg *config.Config, db *sql.DB, pgxPool *pgxpool.Pool) {
	private := r.Group("/api/v1")
	private.Use(jwttoken.Middleware(cfg.JwtSecret))
	private.Use(middleware.ContextMiddleware(db, pgxPool, cfg.Logger))
	private.Use(middleware.AuthMiddleware())
	private.GET("/user/me", auth.Me)
	private.POST("/token/refresh", refreshHandler(cfg.JwtSecret))
}

func loginHandler(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.Login(c, secret)
	}
}

func refreshHandler(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.Refresh(c, secret)
	}
}
