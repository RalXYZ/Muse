package main

import (
	"Muse/conf"
	"Muse/controller"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	conf.Init()

	bot, err := tgbotapi.NewBotAPI(viper.GetString("bot.token"))
	if err != nil {
		logrus.Fatal(err)
	}

	bot.Debug = viper.GetBool("bot.debug")

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.ChannelPost == nil {
			continue
		}
		f := controller.Forward {
			Post: update.ChannelPost,
			Bot: bot,
		}
		if f.NeedsForward() {
			for _, v := range conf.ForwardDest.IdArray {
				v := v
				go f.DoForward(v)
			}
		}
	}
}
