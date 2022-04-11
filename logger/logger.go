package logger

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func NewLogger() *zap.SugaredLogger {
	aa := zap.NewProductionEncoderConfig()
	// fileEncoder := zapcore.NewJSONEncoder(aa)
	aa.EncodeLevel = zapcore.CapitalColorLevelEncoder
	aa.EncodeTime = zapcore.ISO8601TimeEncoder
	aa.MessageKey = `message`
	aa.LevelKey = "severity"
	aa.TimeKey = "time"
	aa.CallerKey = "caller"
	aa.EncodeCaller = zapcore.FullCallerEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(aa)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(colorable.NewColorableStdout()), zap.InfoLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	return logger.Sugar()
}
