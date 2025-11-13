package tgbot

import (
	"fmt"
	"strconv"

	"github.com/PlopyBlopy/notebot/internal/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Token   string `env:"BOT_TOKEN,required"`
	Timeout string `env:"BOT_TIMEOUT"`
}

type Bot struct {
	api    *tgbotapi.BotAPI
	router *router.Router
}

func NewBot(router *router.Router, c Config) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(c.Token)
	if err != nil {
		return nil, fmt.Errorf("couldn't create a new BotAPI. %s", err)
	}

	log.Info().Msgf("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		api:    bot,
		router: router,
	}, nil
}

func (b *Bot) Run(c Config) error {
	u := tgbotapi.NewUpdate(0)

	timeout, err := strconv.Atoi(c.Timeout)
	if err != nil {
		return fmt.Errorf("timeout value invalid. %w", err)
	}

	u.Timeout = timeout

	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		b.router.HandleUpdate(b.api, update)
	}
	return nil
}
