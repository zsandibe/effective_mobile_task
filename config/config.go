package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Driver       string
	User         string
	Password     string
	DatabaseName string
	Host         string
	Port         int
}

func NewConfig() (Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, err
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return Config{}, err
	}

	return Config{
		Driver:       os.Getenv("DRIVER"),
		User:         os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Host:         os.Getenv("DB_HOST"),
		Port:         port,
	}, nil
}
