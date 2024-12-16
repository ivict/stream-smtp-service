package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/ivict/stream-smtp-service/internal/config"
)

func main() {
	// Initialize Zerolog logger with output to stdout
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	cfg := config.LoadConfig()
	zerolog.SetGlobalLevel(cfg)
}
