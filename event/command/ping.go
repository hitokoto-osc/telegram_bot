package command

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

// Ping 查询机器人是否在线
func Ping(b *telebot.Bot) {
	b.Handle("/ping", func(m telebot.Context) error {
		_, err := b.Send(m.Chat(), "Pong!")
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
		}
		return err
	})
}
