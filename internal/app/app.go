package app

import (
	"context"
	"errors"
	"fmt"
	"gocarts/internal/box"
	"gocarts/internal/controllers/http/handlers"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

func Run() error {
	envBox, err := box.New()
	if err != nil {
		return err
	}

	// init service lier

	// init repository lier <- слой

	h := handlers.NewRouter()

	httpServer := &http.Server{
		Handler: h,
		Addr: fmt.Sprintf("%s:%s",
			envBox.Config.HTTPServer.Address,
			envBox.Config.HTTPServer.Port,
		),
	}

	go func() {
		if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("failed to start http server")
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Info().Msg(fmt.Sprintf("listening on %s", envBox.Config.HTTPServer.Address+":"+envBox.Config.HTTPServer.Port))
	<-ctx.Done()

	log.Info().Msg("shutting down server gracefully")

	httpServer.SetKeepAlivesEnabled(false)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("failed to shutdown server")
	}

	if err = envBox.Close(); err != nil {
		log.Error().Err(err).Msg("failed to close connections")
	}

	log.Info().Msg("server gracefully shutdown")

	return nil
}
