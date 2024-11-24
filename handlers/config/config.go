package config

import (
	"os"
)

type Config struct {
	PORT              string
	TELEGRAM_APITOKEN string
}

func New() *Config {
	return &Config{
		PORT:              getEnv("PORT", "8080"),
		TELEGRAM_APITOKEN: getEnv("TELEGRAM_APITOKEN", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
