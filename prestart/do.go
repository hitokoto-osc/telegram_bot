package prestart

import "github.com/hitokoto-osc/telegram_bot/config"

func Do() {
	initConfig(config.Path)
}
