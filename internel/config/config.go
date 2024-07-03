package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppName  string
	Bind     string
	DbConfig string
	Debug    bool
}

func NewConfig() *Config {
	parseBool, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		parseBool = false
	}
	return &Config{
		AppName:  os.Getenv("APP_NAME"),
		Bind:     os.Getenv("BIND"),
		DbConfig: os.Getenv("DB_CONFIG"),
		Debug:    parseBool,
	}
}
