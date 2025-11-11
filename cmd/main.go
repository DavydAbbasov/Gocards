package main

import (
	"gocarts/internal/app"

	log "github.com/rs/zerolog/log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal().Err(err).Msg("Application execution error")
	}
}
