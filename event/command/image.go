package command

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
	"time"
)

// Image 返回随机必应图片
func Image(bot *telebot.Bot) {
	bot.Handle("/image", func(m *telebot.Message) {
		// 暂时先只做 Bing
		photo := &telebot.Photo{
			File: telebot.FromURL("https://uploadbeta.com/api/pictures/random/?key=BingEverydayWallpaperPicture&r=" + strconv.FormatInt(time.Now().UnixNano(), 10)),
		}
		_, err := photo.Send(bot, m.Chat, &telebot.SendOptions{
			ReplyTo: m,
		})
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
		}
	})
}
