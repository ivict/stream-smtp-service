package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	LogLevel     int8   `default:"4" envconfig:"LOG_LEVEL"`
	RedisHost    string `binding:"required" envconfig:"REDIS_HOST"`
	RedisPort    string `binding:"required" envconfig:"REDIS_PORT"`
	SmtpHost     string `binding:"required" envconfig:"SMTP_HOST"`
	SmtpPort     string `binding:"required" envconfig:"SMTP_PORT"`
	SmtpFrom     string `binding:"required" envconfig:"SMTP_FROM"`
	SmtpUser     string `binding:"required" envconfig:"SMTP_USER"`
	SmtpPassword string `binding:"required" envconfig:"SMTP_PASSWORD"`
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		cwd, _ := os.Getwd()
		log.Info().Err(err).Str("cwd", cwd).Msg("Error loading .env file, perhaps running on docker infrastructure.")
	}

	logLevel, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	return Config{
		LogLevel:     int8(logLevel),
		RedisHost:    os.Getenv("REDIS_HOST"),
		RedisPort:    os.Getenv("REDIS_PORT"),
		SmtpHost:     os.Getenv("SMTP_HOST"),
		SmtpPort:     os.Getenv("SMTP_PORT"),
		SmtpFrom:     os.Getenv("SMTP_FROM"),
		SmtpUser:     os.Getenv("SMTP_USER"),
		SmtpPassword: os.Getenv("SMTP_PASSWORD"),
	}
}
