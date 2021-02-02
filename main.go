package main

import (
	"Muse/conf"
	"Muse/controller"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func init() {
	conf.Init()
}

func main() {
	bot, err := tgbotapi.NewBotAPI(viper.GetString("bot.token"))
	if err != nil {
		panic(err)
	}

	bot.Debug = viper.GetBool("bot.debug")

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.ChannelPost == nil {
			continue
		}
		f := controller.Forward{Post: update.ChannelPost, Bot: bot}
		if f.NeedsForward() {
			f.DoForward()
		}
	}
}
