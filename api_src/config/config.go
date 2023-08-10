package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	PORT int

	HASH_SALT string

	DB_DRIVER            string
	DB_CONNECTION_STRING string
}

func NewConfig() *Config {
	return &Config{
		PORT: GetenvInt("PORT", 8080),

		HASH_SALT:            GetenvString("HASH_SALT", "tuc-environment"),
		DB_DRIVER:            GetenvString("DB_DRIVER", "sqlite"),
		DB_CONNECTION_STRING: GetenvString("DB_CONNECTION_STRING", "db.sqlite"),
	}
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
