package telegram

import (
	"gopkg.in/telebot.v3"
	"time"

	"github.com/hitokoto-osc/telegram_bot/config"
	log "github.com/sirupsen/logrus"
)

// InitBot Telegram 机器人的初始化入口
func InitBot() *telebot.Bot {
	c := &config.Telegram{}
	if c.Token() == "" {
		log.Fatal("电报令牌为空。您是不是忘记填写了令牌？")
	}
	bot, err := telebot.NewBot(telebot.Settings{
		URL:     c.Registry(),
		Token:   c.Token(),
		Updates: 100,
		Poller: &telebot.LongPoller{
			Timeout: time.Duration(c.PollInterval()) * time.Second,
		},
		Verbose: false,
		OnError: nil,
		Client:  nil,
		Offline: false,
	})
	if err != nil {
		log.Fatalf("机器人初始化时发生致命错误：\n%s\n", err)
	}
	return bot
}
