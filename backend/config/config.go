package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
<<<<<<< HEAD
	JWTSecret  string
=======
>>>>>>> develop
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
<<<<<<< HEAD
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" || cfg.JWTSecret == "" {
		return nil, fmt.Errorf("missing required environment variables for DB connection or JWT secret")
=======
	}
	if cfg.DBHost == "" || cfg.DBPort == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("missing required environment variables for DB connection")
>>>>>>> develop
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	return cfg, nil
} 
