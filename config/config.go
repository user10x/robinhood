package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Host        string
	DoraHost    string
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

	expires := os.Getenv("EXPIRES_IN")
	expiresIn, _ := strconv.Atoi(expires)
	config := &Config{
		Host:        os.Getenv("HOST"),
		Project:     os.Getenv("PROJECT"),
		Username:    os.Getenv("USERNAME"),
		Password:    os.Getenv("PASSWORD"),
		GrantType:   os.Getenv("GRANT_TYPE"),
		ClientId:    os.Getenv("CLIENT_ID"),
		ExpiresIn:   expiresIn,
		DeviceToken: os.Getenv("DEVICE_TOKEN"),
		Scope:       os.Getenv("SCOPE"),
		DoraHost:    os.Getenv("DORA_HOST"),
	}

	return config
}
