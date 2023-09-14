package config

import (
	"github.com/cockroachdb/errors"
	"github.com/hitokoto-osc/telegram_bot/logging"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
)

// InitConfig 为配置初始化函数
func InitConfig(path string) {
	logger := logging.GetLogger()
	defer logger.Sync()
	logger.Debug("初始化默认配置...")
	loadTelegramConfig()

	logger.Debug("开始读取配置文件...")
	viper.SetEnvPrefix("TELEGRAM_BOT_")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")
	viper.AddConfigPath(".") // 二进制执行目录
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./data")
	if path != "" {
		viper.AddConfigPath(path)
	}
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		var e viper.ConfigFileNotFoundError
		if !errors.As(err, &e) {
			logger.Fatal("无法读取配置文件", zap.Error(err))
		} else {
			logger.Warn("配置文件不存在，使用默认配置（或从环境变量读取）", zap.Error(err))
		}
	}
	logger.Sugar().Debugf("使用配置文件：%s", viper.ConfigFileUsed())
}
