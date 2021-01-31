package main

import (
	"Muse/conf"
	"Muse/controller"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func main() {
	conf.InitConf()
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
		if controller.NeedsForward(update.ChannelPost) {
			controller.DoForward(bot, update.ChannelPost)
		}
	}
}
