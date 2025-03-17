package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	// Import the config package
	_ "github.com/spf13/viper"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Start the server
	if err := StartServer(); err != nil {
		log.Fatal().Err(err).Msg("Could not start server")
	}
}
