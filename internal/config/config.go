package config

import "os"

type Config struct {
	DBUrl     string
	JWTSecret string
	Port      string
}

func Load() *Config {
	return &Config{
		DBUrl:     os.Getenv("DBURL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		Port:      os.Getenv("PORT"),
	}
}
