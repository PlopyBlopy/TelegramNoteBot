package write

import (
	"github.com/PlopyBlopy/notebot/internal/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

// var usecase *Usecase

// type Handler struct {
// 	usecase *Usecase
// }

func NewHandler(usecase *Usecase) router.HandleFunc {
	return func(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
		var response tgbotapi.MessageConfig

		output, err := usecase.WriteNote(update)

		if err != nil {
			log.Error().Err(err).Msg("Handler error.")
			response = tgbotapi.NewMessage(update.Message.Chat.ID, "oops... error...")
		} else {
			response = tgbotapi.NewMessage(update.Message.Chat.ID, output)
		}

		bot.Send(response)
	}
}
