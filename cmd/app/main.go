package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/PlopyBlopy/notebot/config"
	noteservice "github.com/PlopyBlopy/notebot/internal/adapters/note_service"
	"github.com/PlopyBlopy/notebot/internal/router"
	"github.com/PlopyBlopy/notebot/internal/write"
	"github.com/PlopyBlopy/notebot/pkg/logger"
	"github.com/PlopyBlopy/notebot/pkg/tgbot"
	"github.com/rs/zerolog/log"
)

func main() {
	// launch
	logger.NewLogger()
	log.Info().Msg("bot launch")

	c, err := config.InitConfig()
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	if err := App(context.Background(), c); err != nil {
		log.Fatal().Err(err).Msg("failed to initialize application:")
	}
}

func App(ctx context.Context, c config.Config) error {
	store, err := noteservice.NewStore()
	if err != nil {
		return fmt.Errorf("store creation failed: %w", err)
	}

	writeUsecase := write.NewUsecase(store)

	router := router.NewRouter()

	router.HandleCommand("/note", write.NewHandler(writeUsecase))

	// create bot
	bot, err := tgbot.NewBot(router, c.TgBot)
	if err != nil {
		return fmt.Errorf("bot init failed. %w", err)
	}

	// bot started

	errChan := make(chan error, 1)

	go func() {
		log.Info().Msg("bot started")

		err := bot.Run()
		if err != nil {
			errChan <- err
		}
		close(errChan)
	}()

	// Signal for Shutdown or Close
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Info().Msg("A termination signal is received, and the server stops...")
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("bot crashed: %w", err)
		}
	}

	// Close or Shutdown

	log.Info().Msg("bot stopped")

	return nil
}
