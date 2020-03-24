package event

import (
	"gopkg.in/tucnak/telebot.v2"
	"source.hitokoto.cn/hitokoto/telegram_bot/src/event/command"
)

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
