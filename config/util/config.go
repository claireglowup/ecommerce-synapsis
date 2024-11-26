package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver          string
	DatabaseURL       string
	MigrationURL      string
	HTTPServerAddress string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DBDriver:          os.Getenv("DATABASE_DRIVER"),
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		HTTPServerAddress: os.Getenv("HTTP_SERVER_ADDRESS"),
		MigrationURL:      os.Getenv("MIGRATION_URL"),
	}, nil

}
