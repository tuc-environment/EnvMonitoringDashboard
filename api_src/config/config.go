package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	PORT int

	HASH_SALT string

	DB_DRIVER                string
	DB_CONNECTION_STRING     string
	RAW_DB_DRIVER            string
	RAW_DB_CONNECTION_STRING string
}

func NewConfig() *Config {
	godotenv.Load()
	return &Config{
		PORT: GetenvInt("PORT", 8080),

		HASH_SALT:                GetenvString("HASH_SALT", "tuc-environment"),
		DB_DRIVER:                GetenvString("DB_DRIVER", "sqlite"),
		DB_CONNECTION_STRING:     GetenvString("DB_CONNECTION_STRING", "db.sqlite"),
		RAW_DB_DRIVER:            GetenvString("RAW_DB_DRIVER", "mssql"),
		RAW_DB_CONNECTION_STRING: GetenvString("RAW_DB_CONNECTION_STRING", ""),
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
