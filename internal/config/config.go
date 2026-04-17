package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
	//JWTSecret string
	//LogLevel         string
	//HTTPReadTimeout  time.Duration
	//HTTPWriteTimeout time.Duration
}

func mustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		panic(fmt.Sprintf("required env var %s is not set", key))
	}
	return val
}

func getEnv(key, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return def
	}
	return val
}

func Load() Config {
	return Config{
		DatabaseURL: mustGetEnv("DATABASE_URL"),
		Port:        getEnv("PORT", "8080"),
	}
}
