package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Host     string
	Project  string
	Username string
	Password string
}

func NewConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil
	}

	return &Config{
		Host:     os.Getenv("HOST"),
		Project:  os.Getenv("PROJECT"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
}
