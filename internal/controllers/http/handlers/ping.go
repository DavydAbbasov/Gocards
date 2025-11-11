package handlers

import (
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "working as well"); err != nil {
			log.Error().Err(err).Msg("service is not working")
		}
	}
}
