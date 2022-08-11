package main

import (
	"log"
	"telegrambot/bot"
	"telegrambot/client"
	"telegrambot/config"
	"telegrambot/package/repository"
	"telegrambot/package/telegram"
	"telegrambot/server"
	"telegrambot/utils/logrus_log"
)

func main() {

	logrus := logrus_log.GetLogger()
	config := config.Config()
	logrus.Info(config)
	db, err := repository.NewSQLliteDB()
	if err != nil {
		logrus.Fatal("Not Started db :", err)
	}
	logrus.Info("Run DB")
	repo := repository.NewRepository(db)
	server := server.NewServer(repo, logrus)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	client := client.NewClient(config, repo, logrus)
	client.Start()

	botApi := bot.StartBot(config.TelegramAPITOKEN, logrus)

	bot := telegram.NewBot(botApi,  server,client, repo, logrus)
	bot.Start()

	err = bot.Start()

	if err != nil {
		log.Fatal(err)
	}

}
