package command

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
	"time"
)

func Image(bot *telebot.Bot) {
	bot.Handle("/image", func(m *telebot.Message) {
		// 暂时先只做 Bing
		photo := &telebot.Photo{
			File: telebot.FromURL("https://uploadbeta.com/api/pictures/random/?key=BingEverydayWallpaperPicture?r=" + strconv.FormatInt(time.Now().UnixNano(), 10)),
		}
		_, err := bot.Reply(m, photo)
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n")
		}
	})
}
