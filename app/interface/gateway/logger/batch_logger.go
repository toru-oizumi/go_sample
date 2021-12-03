package logger

type BatchLogger interface {
	Debug(message string, err error) error
	Info(message string, err error) error
	Warning(message string, err error) error
	Error(message string, err error) error
	Fatal(message string, err error) error
}
