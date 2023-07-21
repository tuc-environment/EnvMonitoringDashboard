package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func GetenvString(key string, defaults string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaults
	}
	return val
}

func GetenvInt(key string, defaults int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaults
	}
	v, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return defaults
	}
	return int(v)
}
