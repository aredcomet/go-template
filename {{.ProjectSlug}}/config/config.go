package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type StringKey string

func (s StringKey) AsString() string {
	return string(s)
}

const RequestContextKey StringKey = "request_context"

type Config struct {
	PostgresUri string
	ContextKey  StringKey
	JwtSecret   string
	UserIdKey   string
	LogLevel    string
	SqlLogLevel string
	Logger      logrus.FieldLogger
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		PostgresUri: os.Getenv("DATABASE_URL"),
		ContextKey:  RequestContextKey,
		JwtSecret:   os.Getenv("JWT_SECRET"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
		SqlLogLevel: os.Getenv("SQL_LOG_LEVEL"),
		Logger:      CreateLogger(os.Getenv("LOG_LEVEL")),
	}
}

func CreateLogger(logLevel string) logrus.FieldLogger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{})

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logger.WithError(err).WithField("input", logLevel).Warn("Invalid log level")
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	return logger
}
