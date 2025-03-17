package main

import (
	"Cocktail_app/CocktailApp/endpoint"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/spf13/viper"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Register the /read endpoint.
	http.HandleFunc("/read", endpoint.ReadHandler)

	// Start the server using the StartServer function from config.go.
	if err := StartServer(); err != nil {
		log.Fatal().Err(err).Msg("Could not start server")
	}
}
