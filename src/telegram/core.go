package telegram

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"source.hitokoto.cn/hitokoto/telegram_bot/src/config"
	"time"
)

func InitBot() *telebot.Bot {
	config := &config.Telegram{}
	if config.Token() == "" {
		log.Fatal("电报令牌为空。您是不是忘记填写了令牌？")
	}
	bot, err := telebot.NewBot(telebot.Settings{
		URL:     config.Registry(),
		Token:   config.Token(),
		Updates: 0,
		Poller: &telebot.LongPoller{
			Timeout: time.Duration(config.PollInterval()) * time.Second,
		},
		Reporter: nil,
		Client:   nil,
	})
	if err != nil {
		log.Fatalf("机器人初始化时发生致命错误：\n%s\n", err)
	}
	return bot
}
