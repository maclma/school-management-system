package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppEnv     string
	GinMode    string
	ServerPort string

	LogLevel  string
	LogFormat string

	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	DBDriver   string
	DBPath     string

	JWTSecret string
	JWTExpiry int
}

// LoadConfig reads configuration from environment variables with sensible defaults.
func LoadConfig() (*Config, error) {
	cfg := &Config{}

	cfg.AppEnv = getEnv("APP_ENV", "development")
	cfg.GinMode = getEnv("GIN_MODE", "debug")
	cfg.ServerPort = getEnv("SERVER_PORT", "8080")

	cfg.LogLevel = getEnv("LOG_LEVEL", "info")
	cfg.LogFormat = getEnv("LOG_FORMAT", "text")

	cfg.DBHost = getEnv("DB_HOST", "localhost")
	cfg.DBUser = getEnv("DB_USER", "postgres")
	cfg.DBPassword = getEnv("DB_PASSWORD", "postgres")
	cfg.DBName = getEnv("DB_NAME", "school_db")
	cfg.DBPort = getEnv("DB_PORT", "5432")
	cfg.DBSSLMode = getEnv("DB_SSLMODE", "disable")
	cfg.DBDriver = getEnv("DB_DRIVER", "sqlite")
	cfg.DBPath = getEnv("DB_PATH", "school.db")

	cfg.JWTSecret = getEnv("JWT_SECRET", "changeme")
	jwtExpiryStr := getEnv("JWT_EXPIRY", "24")
	jwtExpiry, err := strconv.Atoi(jwtExpiryStr)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRY: %v", err)
	}
	cfg.JWTExpiry = jwtExpiry

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
