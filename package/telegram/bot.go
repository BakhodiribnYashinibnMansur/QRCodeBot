package telegram

import (
	"telegrambot/client"
	"telegrambot/model"
	"telegrambot/package/repository"
	"telegrambot/server"
	"telegrambot/utils/logrus_log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	server  *server.Server
	client  *client.Client
	repo    *repository.Repository
	message model.Message
	logrus  *logrus_log.Logger
}

func NewBot(bot *tgbotapi.BotAPI, server *server.Server, client *client.Client, repo *repository.Repository, logrus *logrus_log.Logger) *Bot {
	return &Bot{
		bot:    bot,
		repo:   repo,
		server: server,
		client: client,
		logrus: logrus,
	}
}

func (bt *Bot) Start() error {
	bot := bt.bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates

			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "WAIT ADMIN ANSWERS")
			// msg.ReplyToMessageID = update.Message.MessageID
			// if _, err := bot.Send(msg); err != nil {
			// 	panic(err)
			// }
			// Handle regular messages
			if err := bt.handleReplyMessage(update.Message); err != nil {
				bt.handleError(update.Message.Chat.ID, err)
			}
			continue
		}
		if update.Message.IsCommand() {
			if err := bt.handleCommand(update.Message); err != nil {
				bt.handleError(update.Message.Chat.ID, err)
			}

			continue
		}

	}
	return nil
}
