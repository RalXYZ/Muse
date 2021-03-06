package controller

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var bot *tgbotapi.BotAPI

func Init() (config tgbotapi.UpdateConfig) {
	logrus.Info("Connecting to Telegram API server...")
	var err error
	bot, err = tgbotapi.NewBotAPI(viper.GetString("bot.token"))
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Telegram API server connected")

	bot.Debug = viper.GetBool("bot.debug")

	config = tgbotapi.NewUpdate(0)
	config.Timeout = 30

	return
}
