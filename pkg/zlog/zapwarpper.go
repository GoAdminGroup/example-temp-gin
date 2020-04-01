package zlog

import (
	"fmt"
	"go.uber.org/zap"
)

var zapLog *zapLogger

type zapLogger struct {
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
}

func S() *zap.SugaredLogger {
	if zapLog == nil {
		panic(fmt.Errorf("not use zlog.NewZapLog()"))
	}
	return zapLog.Sugar
}

func NewZapLog(log *zap.Logger, sugar *zap.SugaredLogger) {
	if zapLog == nil {
		zapLog = &zapLogger{
			Log:   log,
			Sugar: sugar,
		}
		sugar.Warn("zap log init success")
	}
}
