package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

const (
	commandStart = "start"
)

func (bot *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return bot.handleStartCommand(message)
	default:
		return bot.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) (err error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, "YOU PRES START COMMAND")
	_, err = b.bot.Send(msg)
	logrus.Debug(err)
	return err

}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) (err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "WRONG COMMAND")
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) (err error) {

	msg := tgbotapi.NewMessage(message.Chat.ID, "YOU SEND ME MESSAGE,I GOT IT")
	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return err
}
func (b *Bot) handleReplyMessage(message *tgbotapi.Message) (err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "YOU SEND ME MESSAGE,I GOT IT")
	msg.ReplyToMessageID = message.MessageID
	if _, err := b.bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return err
}
func (b *Bot) saveLink(message *tgbotapi.Message, accessToken string) (err error) {
	// if err := b.validateURL(message.Text); err != nil {
	// 	return invalidUrlError
	// }

	// if err := b.client.Add(context.Background(), pocket.AddInput{
	// 	URL:         message.Text,
	// 	AccessToken: accessToken,
	// }); err != nil {
	// 	return unableToSaveError
	// }

	return nil
}
