package zap

import "go.uber.org/zap"

type ZapLogger struct {
	Logger *zap.Logger
}

func NewZapLogger() *ZapLogger {
	logger, _ := zap.NewDevelopment()
	return &ZapLogger{Logger: logger}
}

func (l *ZapLogger) Debug(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	httpStatusCode, _ := args[0].(int)
	message, _ := args[1].(string)

	l.Logger.Debug(
		"InternalServerError",
		zap.Int("status_code", httpStatusCode),
		zap.String("message", message),
	)
	return nil
}

func (l *ZapLogger) Info(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	httpStatusCode, _ := args[0].(int)
	message, _ := args[1].(string)

	logger.Info(
		"InternalServerError",
		zap.Int("status_code", httpStatusCode),
		zap.String("message", message),
	)
	return nil
}

func (l *ZapLogger) Warning(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	httpStatusCode, _ := args[0].(int)
	message, _ := args[1].(string)

	logger.Warn(
		"InternalServerError",
		zap.Int("status_code", httpStatusCode),
		zap.String("message", message),
	)
	return nil
}

func (l *ZapLogger) Error(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	httpStatusCode, _ := args[0].(int)
	message, _ := args[1].(string)

	logger.Error("InternalServerError",
		zap.Int("status_code", httpStatusCode),
		zap.String("message", message),
	)
	return nil
}

func (l *ZapLogger) Fatal(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	message, _ := args[0].(string)

	logger.Fatal("ERROR",
		zap.String("message", message),
	)
	return nil
}
