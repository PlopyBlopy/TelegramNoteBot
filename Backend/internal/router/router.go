package router

import (
	"github.com/PlopyBlopy/notebot/pkg/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type HandleFunc func(*tgbotapi.BotAPI, tgbotapi.Update)

type Router struct {
	handlers map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandleFunc),
	}
}

func (r *Router) HandleCommand(command string, handle HandleFunc) {
	r.handlers[command] = handle
}

func (r *Router) HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	text := update.Message.Text
	if text == "" {
		return
	}

	command, err := message.GetMsgCommand(update)
	if err != nil {
		log.Error().Err(err).Msgf("command error.")
		return
	}

	if handler, ok := r.handlers[command]; ok {
		handler(bot, update)
		return
	}

	sendMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "The command was not found")
	bot.Send(sendMsg)
}
