package command

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func Ping(b *telebot.Bot) {
	b.Handle("/ping", func(m *telebot.Message) {
		_, err := b.Send(m.Chat, "Pong!")
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
		}
	})
}
