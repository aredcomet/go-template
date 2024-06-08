package database

import (
	"context"
	"database/sql"

	pgxlogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/sirupsen/logrus"
)

func Connect(dsn string, logLevel string, fieldLogger logrus.FieldLogger) *sql.DB {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fieldLogger.Fatalf("unable to connect to database: %v", err)
	}
	var greeting string
	err = db.QueryRow("SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fieldLogger.Fatalf("QueryRow failed: %v\n", err)
	}
	fieldLogger.Info(greeting)
	return db
}

func ConnectPool(dsn string, logLevel string, fieldLogger logrus.FieldLogger) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		fieldLogger.Fatalf("Failed to parse dsn config: %v", err)
	}
	traceLogLevel, err := tracelog.LogLevelFromString(logLevel)
	if err != nil {
		fieldLogger.Fatalf("Failed to set tracelog: %v", err)
	}
	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxlogrus.NewLogger(fieldLogger),
		LogLevel: traceLogLevel,
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fieldLogger.Fatalf("Failed connect: %v", err)
	}

	var greeting string
	err = pool.QueryRow(context.Background(), "select 'Connected to pgx pool...!'").Scan(&greeting)
	if err != nil {
		fieldLogger.Fatalf("Failed to do initial query: %v", err)
	}

	fieldLogger.Info(greeting)
	return pool
}
