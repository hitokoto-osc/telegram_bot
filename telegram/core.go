package telegram

import (
	"encoding/json"
	"github.com/hitokoto-osc/telegram_bot/request"
	"go.uber.org/zap"
	"gopkg.in/telebot.v3"
	"time"

	"github.com/hitokoto-osc/telegram_bot/config"
)

// InitBot Telegram 机器人的初始化入口
func InitBot() *telebot.Bot {
	defer zap.L().Sync()
	c := &config.Telegram{}
	if c.Token() == "" {
		zap.L().Fatal("电报令牌为空。您是不是忘记填写了令牌？")
	}
	bot, err := telebot.NewBot(telebot.Settings{
		URL:     c.Registry(),
		Token:   c.Token(),
		Updates: 100,
		Poller: &telebot.LongPoller{
			Timeout: time.Duration(c.PollInterval()) * time.Second,
		},
		Verbose: config.Debug,
		OnError: func(err error, ctx telebot.Context) {
			if ctx == nil {
				zap.L().Error("机器人发生未知错误。", zap.Error(err))
				return
			}
			zap.L().Error("处理过程中发生错误。", zap.Error(err), zap.String("context", ctx.Text()))
			_ = ctx.Reply("处理过程中发生错误，请稍后再试！")
		},
		Client:  request.NewWithCallerSkip(7).StandardClient(),
		Offline: false,
	})
	if err != nil {
		zap.L().Fatal("机器人初始化时发生致命错误。", zap.Error(err))
	}
	bot.Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(ctx telebot.Context) error {
			data, _ := json.MarshalIndent(ctx.Update(), "", "  ")
			zap.L().Debug("收到消息", zap.ByteString("data", data))
			return next(ctx)
		}
	})
	return bot
}
