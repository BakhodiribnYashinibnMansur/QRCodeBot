package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleError(chatID int64, err error) {
	var messageText string

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}
