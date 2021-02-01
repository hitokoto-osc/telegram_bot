package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitConfig 为配置初始化函数
func InitConfig(path string) {
	log.Debug("初始化默认配置...")
	loadTelegramConfig()

	log.Debug("开始读取配置文件...")
	viper.SetConfigName("config")
	viper.AddConfigPath(".") // 二进制执行目录
	if path != "" {
		viper.AddConfigPath(path)
	}
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err) // 读取配置文件失败致命错误
	}
	log.Debugln("使用配置文件：", viper.ConfigFileUsed())
}
