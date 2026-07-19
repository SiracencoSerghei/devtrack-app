package config

import (
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	Port        string
	JWTSecret   string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	CORSOrigin  string
}

func LoadConfig() *Config {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" || secret == "required_secret_key_change_me_in_production" {
		slog.Error("CRITICAL: JWT_SECRET environment variable is missing or insecure!")
		os.Exit(1)
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		JWTSecret:   secret,
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "password"),
		DBName:      getEnv("DB_NAME", "devtrack"),
		CORSOrigin:  getEnv("CORS_ORIGIN", "http://localhost:5173"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.TrimSpace(value)
	}
	return defaultValue
}