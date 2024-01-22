package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database databaseConfig
	Server   serverConfig
	Api      apiConfig
}

type databaseConfig struct {
	Driver       string
	User         string
	Password     string
	DatabaseName string
	Host         string
	Port         int
}

type serverConfig struct {
	Host string
	Port string
}
type apiConfig struct {
	AgeURL         string
	GenderURL      string
	NationalityURL string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		databaseConfig{
			Driver:       os.Getenv("DRIVER"),
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			DatabaseName: os.Getenv("DB_NAME"),
			Host:         os.Getenv("DB_HOST"),
			Port:         port,
		},
		serverConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		apiConfig{
			AgeURL:         os.Getenv("AGE_API"),
			GenderURL:      os.Getenv("GENDER_API"),
			NationalityURL: os.Getenv("NATIONALITY_API"),
		},
	}, nil
}
