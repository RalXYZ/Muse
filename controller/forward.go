package controller

import (
	"Muse/conf"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type message struct {
	tgbotapi.Message
}

func (m *message) needsFwd() bool {
	for _, v := range conf.ForwardSrc.IdArray {
		v := v
		if v == m.Chat.ID {
			return true
		}
	}
	return false
}

func (m *message) syncFwd() {
	for _, v := range conf.ForwardDest.IdArray {
		go func(chatID int64) {
			msg := m.createSyncMsg(chatID)
			_, err := bot.Send(msg)
			if err != nil {
				logrus.Error(err)
			}
		}(v)
	}
}

func (m *message) createSyncMsg(chatID int64) tgbotapi.Chattable {
	if m.Text != "" {
		config := tgbotapi.NewMessage(chatID, m.Text)
		return config
	} else if m.Photo != nil {
		config := tgbotapi.NewPhotoShare(chatID, m.Photo[0].FileID)
		config.Caption = m.Caption
		return config
	} else if m.Document != nil {
		config := tgbotapi.NewDocumentShare(chatID, m.Document.FileID)
		config.Caption = m.Caption
		return config
	} else if m.Video != nil {
		config := tgbotapi.NewVideoShare(chatID, m.Video.FileID)
		config.Caption = m.Caption
		return config
	} else if m.Voice != nil {
		config := tgbotapi.NewVoiceShare(chatID, m.Voice.FileID)
		config.Caption = m.Caption
		return config
	}
	return nil
}
