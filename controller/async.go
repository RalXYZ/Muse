package controller

import (
	"Muse/conf"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type asyncFwd struct {
	sync.Mutex
	m map[string] []interface{}
}

var asyncFwdMap = asyncFwd{m: make(map[string] []interface{})}

func (m *message) asyncFwd() {
	media := m.mediaGroupProc()
	if media == nil {
		return
	}
	for _, v := range conf.ForwardDest.IdArray {
		go func(chatID int64) {
			msg := tgbotapi.NewMediaGroup(chatID, media)
			_, err := bot.Send(msg)
			if err != nil {
				logrus.Error(err)
			}
		}(v)
	}
}

func (m *message) mediaGroupProc() []interface{} {
	if m.Photo != nil {
		photo := tgbotapi.NewInputMediaPhoto(m.Photo[0].FileID)
		photo.Caption = m.Caption
		asyncFwdMap.makeOrAppend(m.MediaGroupID, photo)
	} else if m.Video != nil {
		video := tgbotapi.NewInputMediaVideo(m.Video.FileID)
		video.Caption = m.Caption
		asyncFwdMap.makeOrAppend(m.MediaGroupID, video)
	}

	time.Sleep(1 * time.Second)
	return asyncFwdMap.popOrSkip(m.MediaGroupID)
}

func (a *asyncFwd) makeOrAppend(mediaGroupID string, media interface{}) {
	a.Lock()
	defer a.Unlock()
	if val, ok := a.m[mediaGroupID]; !ok {
		a.m[mediaGroupID] = make([]interface{}, 1)
		a.m[mediaGroupID][0] = media
	} else {
		val = append(val, media)
		a.m[mediaGroupID] = val
	}
}

func (a *asyncFwd) popOrSkip(mediaGroupID string) []interface{} {
	a.Lock()
	defer a.Unlock()
	if val, ok := a.m[mediaGroupID]; ok {
		delete(a.m, mediaGroupID)
		return val
	} else {
		return nil
	}
}
