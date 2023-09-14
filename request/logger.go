package request

import "go.uber.org/zap"

type loggerWrapper struct {
	l *zap.Logger
}

func newLoggerWrapper(l *zap.Logger) *loggerWrapper {
	return &loggerWrapper{l: l}
}

func (l *loggerWrapper) Error(msg string, keysAndValues ...interface{}) {
	defer l.l.Sync()
	l.l.Sugar().Errorw(msg, keysAndValues...)
}

func (l *loggerWrapper) Info(msg string, keysAndValues ...interface{}) {
	defer l.l.Sync()
	l.l.Sugar().Infow(msg, keysAndValues...)
}

func (l *loggerWrapper) Debug(msg string, keysAndValues ...interface{}) {
	defer l.l.Sync()
	l.l.Sugar().Debugw(msg, keysAndValues...)
}

func (l *loggerWrapper) Warn(msg string, keysAndValues ...interface{}) {
	defer l.l.Sync()
	l.l.Sugar().Warnw(msg, keysAndValues...)
}
