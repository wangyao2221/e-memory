package components

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"e-memory/configs"
)

func NewAccessLogger(config configs.Config) *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.LogConfig.AccessLogPath,
		MaxSize:    config.LogConfig.MaxSizeMB,
		MaxBackups: config.LogConfig.MaxBackupSize,
		MaxAge:     config.LogConfig.MaxAge,
	})

	logConfig := zap.NewProductionEncoderConfig()
	logConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(logConfig), w, zap.InfoLevel)
	return zap.New(core)
}
