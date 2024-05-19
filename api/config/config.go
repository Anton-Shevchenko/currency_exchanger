package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServicePort   string
	MailerHost    string
	ExchangerHost string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		ServicePort:   os.Getenv("API_SERVICE_PORT"),
		MailerHost:    os.Getenv("MAILER_HOST"),
		ExchangerHost: os.Getenv("EXCHANGER_HOST"),
	}

	return config, nil
}
