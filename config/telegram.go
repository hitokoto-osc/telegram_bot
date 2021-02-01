package config

import "github.com/spf13/viper"

func loadTelegramConfig() {
	viper.SetDefault("telegram.token", "")
	viper.SetDefault("telegram.registry", "https://api.telegram.org")
	viper.SetDefault("telegram.poll_interval", 1)
}

// Telegram 定义了 Telegram 初始化所需要的配置
type Telegram struct{}

// Token 为机器人令牌
func (p *Telegram) Token() string {
	return viper.GetString("telegram.token")
}

// Registry 为 Telegram Bot 的通讯服务器
func (p *Telegram) Registry() string {
	return viper.GetString("telegram.registry")
}

// PollInterval 为拉取 Telegram 消息的间隔
func (p *Telegram) PollInterval() uint32 {
	return viper.GetUint32("telegram.poll_interval")
}
