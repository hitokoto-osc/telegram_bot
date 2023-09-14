package logging

import (
	"github.com/cockroachdb/errors"
	"github.com/hitokoto-osc/telegram_bot/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	defer setZapGlobalLogger()
	var err error
	logger, err = build(config.Debug)
	if err != nil {
		panic(errors.Wrap(err, "can't initialize logger"))
	}
}

func buildConfig(isDebug bool) zap.Config {
	var c zap.Config
	if isDebug {
		c = zap.NewDevelopmentConfig() // Console Encoding
		c.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	} else {
		c = zap.NewProductionConfig() // JSON Encoding
		c.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	// 统一配置
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	c.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	c.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder

	c.OutputPaths = []string{"stdout"}
	c.ErrorOutputPaths = []string{"stderr"}

	return c
}

func build(isDebug bool) (*zap.Logger, error) {
	instance, err := buildConfig(isDebug).Build(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	if err != nil {
		return nil, errors.Wrap(err, "can't initialize logger")
	}
	defer instance.Sync()
	return instance, nil
}
