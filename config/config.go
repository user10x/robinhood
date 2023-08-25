package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Host        string
	Project     string
	Username    string
	Password    string
	GrantType   string
	ClientId    string
	ExpiresIn   int
	DeviceToken string
	Scope       string
}

func NewConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	expires := os.Getenv("EXPIRESIN")
	expiresIn, _ := strconv.Atoi(expires)

	return &Config{
		Host:        os.Getenv("HOST"),
		Project:     os.Getenv("PROJECT"),
		Username:    os.Getenv("USERNAME"),
		Password:    os.Getenv("PASSWORD"),
		GrantType:   os.Getenv("PASSWORD"),
		ClientId:    os.Getenv("PASSWORD"),
		ExpiresIn:   expiresIn,
		DeviceToken: os.Getenv("PASSWORD"),
		Scope:       os.Getenv("PASSWORD"),
	}
}
