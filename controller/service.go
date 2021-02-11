package controller

import (
	"Muse/conf"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartService(config tgbotapi.UpdateConfig) {
	updates := bot.GetUpdatesChan(config)

	for update := range updates {
		if update.ChannelPost == nil {
			continue
		} else if !needFwd(update.ChannelPost.Chat.ID) {
			continue
		}

		srcMsg := message{Message: *update.ChannelPost}
		if srcMsg.MediaGroupID == "" {
			srcMsg.syncFwd()
		} else {
			go srcMsg.asyncFwd()
		}
	}
}

func needFwd(chatID int64) bool {
	_, ok := conf.FwdRuleMap[chatID]
	return ok
}
