package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	Port         string
	JWTSecret    string
	S3Region     string
	S3BucketName string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		// DBHost:       os.Getenv("DB_HOST"),
		// DBPort:       os.Getenv("DB_PORT"),
		// DBUser:       os.Getenv("DB_USER"),
		// DBPassword:   os.Getenv("DB_PASSWORD"),
		// DBName:       os.Getenv("DB_NAME"),
		Port:         os.Getenv("PORT"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		S3Region:     os.Getenv("S3_REGION"),
		S3BucketName: os.Getenv("S3_BUCKET_NAME"),
	}
	if cfg.JWTSecret == "" || cfg.S3Region == "" || cfg.S3BucketName == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	return cfg, nil
}
