package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost                 string
	DBPort                 string
	DBUser                 string
	DBPassword             string
	DBName                 string
	Port                   string
	JWTSecret              string
	StorageProvider        string
	S3Region               string
	S3BucketName           string
	SupabaseURL            string
	SupabaseBucket         string
	SupabaseServiceRoleKey string
}

func LoadConfig() (*Config, error) {
	provider := os.Getenv("STORAGE_PROVIDER")
	if provider == "" {
		provider = "s3"
	}

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

	switch cfg.StorageProvider {
	case "s3":
		if cfg.S3Region == "" || cfg.S3BucketName == "" {
			return nil, fmt.Errorf("missing required environment variables for S3 configuration")
		}
	case "supabase":
		if cfg.SupabaseURL == "" || cfg.SupabaseBucket == "" || cfg.SupabaseServiceRoleKey == "" {
			return nil, fmt.Errorf("missing required environment variables for Supabase storage configuration")
		}
	default:
		return nil, fmt.Errorf("unsupported storage provider: %s", cfg.StorageProvider)
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg, nil
}
