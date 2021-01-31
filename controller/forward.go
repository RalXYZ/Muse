package controller

import (
	"Muse/conf"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NeedsForward(message *tgbotapi.Message) bool {
	for _, v := range conf.ForwardSrc.IdArray {
		v := v
		fmt.Println(message)
		if v == message.Chat.ID {
			return true
		}
	}
	for _, v := range conf.ForwardSrc.UserNameArray {
		v := v
		if v == message.Chat.UserName {
			return true
		}
	}

	return false
}

func DoForward(bot *tgbotapi.BotAPI, message *tgbotapi.Message) bool {
	for _, v := range conf.ForwardDest.IdArray {
		v := v
		msg := tgbotapi.NewMessage(v, message.Text)
		_, err := bot.Send(msg)
		if err != nil {
			panic(err)
		}
	}
	for _, v := range conf.ForwardDest.UserNameArray {
		v := v
		msg := tgbotapi.NewMessage(0, message.Text)
		msg.ChannelUsername = v
		_, err := bot.Send(msg)
		if err != nil {
			panic(err)
		}
	}

	return false
}
