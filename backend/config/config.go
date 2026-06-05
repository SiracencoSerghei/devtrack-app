package config

import (
	"os"
	"fmt"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

func LoadConfig() *Config {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default_local_secret_key"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASSWORD")
		name := os.Getenv("DB_NAME")
		
		if host == "" { host = "localhost" }
		if user == "" { user = "postgres" }
		if pass == "" { pass = "password" }
		if name == "" { name = "devtrack" }

		dbURL = fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, pass, host, name)
	}

	return &Config{
		Port:        port,
		JWTSecret:   jwtSecret,
		DatabaseURL: dbURL,
	}
}