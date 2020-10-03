package command

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

func Start(bot *telebot.Bot) {
	bot.Handle("/start", func(m *telebot.Message) {
		if !m.Private() { // 如果不是私发消息，不回复
			return
		}
		_, err := bot.Send(m.Sender, `欢迎您选用一言的服务。
在这里，你可以领略天之高，地之深，可以感受思维的边疆，领略美好，收获感动。还在犹豫什么，快来和我们一起玩耍吧！

你可以...
使用 /about 以更深入得了解我，
使用 /help  查看使用机器人的说明`)
		if err != nil {
			log.Errorf("发送消息时发生了错误，错误信息： %s \n", err)
		}
	})
}
