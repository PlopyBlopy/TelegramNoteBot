package write

import (
	"fmt"
	"time"

	noteservice "github.com/PlopyBlopy/notebot/internal/adapters/note_service"
	"github.com/PlopyBlopy/notebot/pkg/message"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
)

type NoteService interface {
	Write(name string, note noteservice.Note) error
}

type Usecase struct {
	store NoteService
}

func NewUsecase(store NoteService) *Usecase {
	return &Usecase{
		store: store,
	}
}

func (u *Usecase) WriteNote(update tgbotapi.Update) (string, error) {
	c, err := message.GetCommand(update)
	if err != nil {
		return "", fmt.Errorf("failed get command from message in WriteNote. %w", err)
	}

	msg, err := message.GetMsgText(update)
	if err != nil {
		return "", fmt.Errorf("failed get msg from message in WriteNote. %w", err)
	}

	note := noteservice.Note{
		Id:        uuid.New(),
		CreatedAt: time.Now(),
		Text:      msg,
	}

	err = u.store.Write(c, note)
	if err != nil {
		return "", fmt.Errorf("failed to write note. %w", err)
	}

	return "writed", nil
}
