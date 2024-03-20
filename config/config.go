package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort      string
	ClientId     string
	ClientSecret string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type Config struct {
	ApiConfig
	DbConfig
}

func (c *Config) readConfig() error {
	err := godotenv.Load()
	if err != nil {
		panic("Error Loading .env file")
	}

	c.ApiConfig = ApiConfig{
		ApiPort:      os.Getenv("API_PORT"),
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("PASSWORD"),
		User:     os.Getenv("USER"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	if c.ApiConfig.ApiPort == "" || c.ApiConfig.ClientId == "" || c.ApiConfig.ClientSecret == "" || c.DbConfig.Driver == "" || c.DbConfig.Host == "" || c.DbConfig.Name == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" {
		return errors.New("all environment Required")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}
