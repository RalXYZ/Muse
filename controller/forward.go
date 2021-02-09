package controller

import (
	"Muse/conf"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Forward struct {
	Post *tgbotapi.Message
	Bot  *tgbotapi.BotAPI
}

func (f *Forward) NeedsForward() bool {
	for _, v := range conf.ForwardSrc.IdArray {
		v := v
		if v == f.Post.Chat.ID {
			return true
		}
	}

	return false
}

func (f *Forward) DoForward(chatID int64) {
	msg := f.createMessage(chatID)
	_, err := f.Bot.Send(msg)
	if err != nil {
		logrus.Error(err)
	}
}

func (f *Forward) createMessage(chatID int64) tgbotapi.Chattable {
	if f.Post.Text != "" {
		config := tgbotapi.NewMessage(chatID, f.Post.Text)
		return config
	} else if f.Post.Photo != nil {
		if photoArrLength := len(f.Post.Photo); photoArrLength == 1 {
			config := tgbotapi.NewPhotoShare(chatID, f.Post.Photo[0].FileID)
			config.Caption = f.Post.Caption
			return config
		} else {  // FIXME: find a better way dealing with media group
			var photoFiles []interface{}
			for k, v := range f.Post.Photo {
				photoElement := tgbotapi.NewInputMediaPhoto(v.FileID)
				if k == 0 {
					photoElement.Caption = f.Post.Caption
				}
				photoFiles = append(photoFiles, )
			}
			config := tgbotapi.NewMediaGroup(chatID, photoFiles)
			return config
		}
	} else if f.Post.Document != nil {
		config := tgbotapi.NewDocumentShare(chatID, f.Post.Document.FileID)
		config.Caption = f.Post.Caption
		return config
	} else if f.Post.Video != nil {
		config := tgbotapi.NewVideoShare(chatID, f.Post.Video.FileID)
		config.Caption = f.Post.Caption
		return config
	} else if f.Post.Voice != nil {
		config := tgbotapi.NewVoiceShare(chatID, f.Post.Voice.FileID)
		config.Caption = f.Post.Caption
		return config
	}
	return nil
}
