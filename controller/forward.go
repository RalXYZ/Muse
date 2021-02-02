package controller

import (
	"Muse/conf"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Forward struct {
	Post *tgbotapi.Message
	Bot  *tgbotapi.BotAPI
}

func (f *Forward) NeedsForward() bool {
	for _, v := range conf.ForwardSrc.IdArray {
		v := v
		fmt.Println(f.Post)
		if v == f.Post.Chat.ID {
			return true
		}
	}
	for _, v := range conf.ForwardSrc.UserNameArray {
		v := v
		if v == f.Post.Chat.UserName {
			return true
		}
	}

	return false
}

func (f *Forward) DoForward() bool {
	for _, v := range conf.ForwardDest.IdArray {
		v := v
		msg := tgbotapi.NewMessage(v, f.Post.Text)
		_, err := f.Bot.Send(msg)
		if err != nil {
			panic(err)
		}
	}
	for _, v := range conf.ForwardDest.UserNameArray {
		v := v
		msg := tgbotapi.NewMessage(0, f.Post.Text)
		msg.ChannelUsername = v
		_, err := f.Bot.Send(msg)
		if err != nil {
			panic(err)
		}
	}

	return false
}
