package config

import "github.com/spf13/viper"

func loadTelegramConfig() {
	viper.SetDefault("telegram.token", "")
	viper.SetDefault("telegram.registry", "https://api.telegram.org")
	viper.SetDefault("telegram.poll_interval", 1)
}

type Telegram struct{}

func (p *Telegram) Token() string {
	return viper.GetString("telegram.token")
}

func (p *Telegram) Registry() string {
	return viper.GetString("telegram.registry")
}

func (p *Telegram) PollInterval() uint32 {
	return viper.GetUint32("telegram.poll_interval")
}
