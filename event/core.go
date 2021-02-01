package event

import (
	"github.com/hitokoto-osc/telegram_bot/event/command"
	"gopkg.in/tucnak/telebot.v2"
)

// RegisterEvent 定义了机器人事件处理的入口
func RegisterEvent(bot *telebot.Bot) {
	registerCommand(bot)
}

func registerCommand(bot *telebot.Bot) {
	command.Hitokoto(bot)
	command.Image(bot)
	command.Start(bot)
	command.About(bot)
	command.Ping(bot)
	command.Help(bot)
	command.Status(bot)
}
