package command

import (
	"gopkg.in/telebot.v3"
)

// Ping 查询机器人是否在线
func Ping(b *telebot.Bot) {
	b.Handle("/ping", func(m telebot.Context) error {
		_, err := b.Send(m.Chat(), "Pong!")
		return err
	})
}
