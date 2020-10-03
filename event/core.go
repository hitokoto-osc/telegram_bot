package event

import (
	"github.com/hitokoto-osc/telegram_bot/event/command"
	"gopkg.in/tucnak/telebot.v2"
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
