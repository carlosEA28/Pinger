package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Worker   WorkerConfig
}

type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	URL      string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type WorkerConfig struct {
	SchedulerInterval time.Duration
	HTTPTimeout       time.Duration
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			URL:      getEnv("DATABASE_URL", ""),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "pinger"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		Worker: WorkerConfig{
			SchedulerInterval: getEnvDurationSeconds("SCHEDULER_INTERVAL_SECONDS", 30),
			HTTPTimeout:       getEnvDurationSeconds("PING_HTTP_TIMEOUT_SECONDS", 10),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvDurationSeconds(key string, defaultValue int) time.Duration {
	value := getEnv(key, "")
	if value == "" {
		return time.Duration(defaultValue) * time.Second
	}

	seconds, err := strconv.Atoi(value)
	if err != nil || seconds <= 0 {
		return time.Duration(defaultValue) * time.Second
	}

	return time.Duration(seconds) * time.Second
}
