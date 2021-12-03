package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapBatchLogger struct {
	Logger *zap.Logger
}

func NewZapBatchLogger() *ZapBatchLogger {
	logger, _ := zap.NewDevelopment()
	return &ZapBatchLogger{Logger: logger}
}

func (l *ZapBatchLogger) Debug(message string, err error) error {
	logger := l.Logger
	defer logger.Sync()
	return logging(logger.Debug, message, err)
}

func (l *ZapBatchLogger) Info(message string, err error) error {
	logger := l.Logger
	defer logger.Sync()
	return logging(logger.Info, message, err)
}

func (l *ZapBatchLogger) Warning(message string, err error) error {
	logger := l.Logger
	defer logger.Sync()
	return logging(logger.Warn, message, err)
}

func (l *ZapBatchLogger) Error(message string, err error) error {
	logger := l.Logger
	defer logger.Sync()
	return logging(logger.Error, message, err)
}

func (l *ZapBatchLogger) Fatal(message string, err error) error {
	logger := l.Logger
	defer logger.Sync()
	return logging(logger.Fatal, message, err)
}

func logging(logger func(msg string, fields ...zapcore.Field), message string, err error) error {
	if err != nil {
		logger(
			message,
			zap.String("error", err.Error()),
		)
	} else {
		logger(
			message,
		)
	}
	return err
}
