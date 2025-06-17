package config

import "os"

type Config struct {
	DBUrl     string
	JWTSecret string
	Port      string
}

func Load() *Config {
	return &Config{
		DBUrl:     getEnv("DBURL"),
		JWTSecret: getEnv("JWT_SECRET"),
		Port:      getEnv("PORT"),
	}
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return ""
}
