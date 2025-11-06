package message

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// get command with format: / - only slash, /n - (without space), /n_ - (with space)
func GetMsgCommand(update tgbotapi.Update) (string, error) {
	text := update.Message.Text

	if len(text) == 0 {
		return "", fmt.Errorf("text is empty")
	}

	if !strings.HasPrefix(text, "/") {
		return "", fmt.Errorf("not have command slash")
	}

	idx := getSeparatorIndex(text)
	if idx == -1 {
		return text, nil
	}

	command := text[:idx]

	return command, nil
}

// only command: some - command without slash
func GetCommand(update tgbotapi.Update) (string, error) {
	text := update.Message.Text

	if len(text) == 0 {
		return "", fmt.Errorf("text is empty")
	}

	if !strings.HasPrefix(text, "/") {
		return "", fmt.Errorf("not have command slash")
	}

	idx := getSeparatorIndex(text)
	if idx == -1 {
		return text, nil
	}

	command := text[1:idx]

	return command, nil
}

// can get message text if have command prefix
func GetMsgText(update tgbotapi.Update) (string, error) {
	text := update.Message.Text

	// if command empty - can't process message text
	_, err := GetMsgCommand(update)
	if err != nil {
		return "", fmt.Errorf("get message command error. %w", err)
	}

	idx := getSeparatorIndex(text)
	if idx == -1 {
		return "", fmt.Errorf("message empty. %w", err)
	}

	textContent := text[idx+1:]
	if len(textContent) == 0 {
		return "", fmt.Errorf("message empty")
	}

	return textContent, nil
}

// get space separator index for command_text
func getSeparatorIndex(text string) int {
	return strings.Index(text, " ")
}
