package main

import (
	"github.com/hitokoto-osc/telegram_bot/config"
	"github.com/hitokoto-osc/telegram_bot/event"
	"github.com/hitokoto-osc/telegram_bot/flags"
	"github.com/hitokoto-osc/telegram_bot/telegram"
	"go.uber.org/zap"
)

func init() {
	flags.Do()
	config.InitConfig(config.Path)
}

func main() {
	zap.L().Info("开始初始化机器人...")
	bot := telegram.InitBot()
	// 注册机器人事件
	event.RegisterEvent(bot)
	go bot.Start()
	zap.L().Info("机器人初始化完成，开始接收信息。")
	select {} // revive:disable-line:empty-block 错误检测。 堵塞进程是必须的
}
