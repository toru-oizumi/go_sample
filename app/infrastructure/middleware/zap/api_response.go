package zap

import "go.uber.org/zap"

type ZapApiResponseLogger struct {
	Logger *zap.Logger
}

func NewZapApiResponseLogger() *ZapApiResponseLogger {
	logger, _ := zap.NewDevelopment()
	return &ZapApiResponseLogger{Logger: logger}
}

func (l *ZapApiResponseLogger) Debug(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	statusCode, _ := args[0].(int)
	message, _ := args[1].(string)

	l.Logger.Debug(
		"API Response",
		zap.Int("status_code", statusCode),
		zap.String("message", message),
	)
	return nil
}

func (l *ZapApiResponseLogger) Info(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	statusCode, _ := args[0].(int)
	message, _ := args[1].(string)

	logger.Info(
		"API Response",
		zap.Int("status_code", statusCode),
		zap.String("message", message),
	)
	return nil
}

func (l *ZapApiResponseLogger) Warning(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	statusCode, _ := args[0].(int)
	message, _ := args[1].(string)

	logger.Warn(
		"API Response",
		zap.Int("status_code", statusCode),
		zap.String("message", message),
	)
	return nil
}

func (l *ZapApiResponseLogger) Error(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	statusCode, _ := args[0].(int)
	message, _ := args[1].(string)

	logger.Error(
		"API Response",
		zap.Int("status_code", statusCode),
		zap.String("message", message),
	)
	return nil
}

func (l *ZapApiResponseLogger) Fatal(args ...interface{}) error {
	logger := l.Logger
	defer logger.Sync()

	statusCode, _ := args[0].(int)
	message, _ := args[1].(string)

	logger.Fatal(
		"API Response",
		zap.Int("status_code", statusCode),
		zap.String("message", message),
	)
	return nil
}
