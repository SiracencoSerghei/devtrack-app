package config

import (
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

// LoadConfig зчитує системні змінні оточення напряму
func LoadConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		JWTSecret:   getEnv("JWT_SECRET", "required_secret_key_change_me_in_production"),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "password"),
		DBName:      getEnv("DB_NAME", "devtrack"),
		CORSOrigin:  getEnv("CORS_ORIGIN", "http://localhost:5173"),
	}
}

// Допоміжна функція для дефолтних значень
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.TrimSpace(value)
	}
	return defaultValue
}