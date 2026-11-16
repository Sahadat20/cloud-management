package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "ec2idp"),
			Password: getEnv("DB_PASSWORD", "ec2idp_secret"),
			DBName:   getEnv("DB_NAME", "ec2idp"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}, nil
}
func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.DBName, d.SSLMode,
	)
}

func getEnv(key, defaulvalue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaulvalue
}
