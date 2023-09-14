package command

import (
	"gopkg.in/telebot.v3"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// Image 返回随机必应图片
func Image(bot *telebot.Bot) {
	bot.Handle("/image", func(ctx telebot.Context) error {
		// 暂时先只做 Bing
		photo := &telebot.Photo{
			File: telebot.FromURL("https://uploadbeta.com/api/pictures/random/?key=BingEverydayWallpaperPicture&r=" + strconv.FormatInt(time.Now().UnixNano(), 10)),
		}
		_, err := photo.Send(bot, ctx.Chat(), &telebot.SendOptions{
			ReplyTo: ctx.Message(),
		})
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
			return err
		}
		return nil
	})
}
