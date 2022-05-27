package config

import (
	"github.com/joho/godotenv"
)

// Load function loads config from file
func Load() error {
	err := godotenv.Load(".env")
	return err
}
