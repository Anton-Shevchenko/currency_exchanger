package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServicePort      string
	ExchangerHost    string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	MailSender       string
	MailPrivKey      string
	MailPubKey       string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		ServicePort:      os.Getenv("MAILER_SERVICE_PORT"),
		ExchangerHost:    os.Getenv("EXCHANGER_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USERNAME"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		MailSender:       os.Getenv("SENDER_MAIL"),
		MailPrivKey:      os.Getenv("PRIV_KEY_MAIL"),
		MailPubKey:       os.Getenv("PUB_KEY_MAIL"),
	}

	return config, nil
}
