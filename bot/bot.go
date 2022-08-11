package bot

import (
	"telegrambot/utils/logrus_log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string, logrus *logrus_log.Logger) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logrus.Panic(err)
	}
	bot.Debug = true

	logrus.Printf("Authorized on account %s", bot.Self.UserName)
	logrus.Info("Start Telegram Bot")

	return bot
}
