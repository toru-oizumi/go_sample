package logger

type WsApiLogger interface {
	Debug(...interface{}) error
	Info(...interface{}) error
	Warning(...interface{}) error
	Error(...interface{}) error
	Fatal(...interface{}) error
}
