package handler

import (
	"go_sample/app/domain/repository"
	"go_sample/app/interface/gateway/logger"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

type WsHandler struct {
	connection repository.Connection
	logger     logger.WsApiLogger
}

func NewWsHandler(
	connection repository.Connection,
	logger logger.WsApiLogger,
) *WsHandler {
	return &WsHandler{
		connection: connection,
		logger:     logger,
	}
}

func (h *WsHandler) Chat() *ChatWsHandler {
	return &ChatWsHandler{
		Logger: h.logger,
	}
}

func (h *WsHandler) Play() *PlayWsHandler {
	return &PlayWsHandler{
		Logger: h.logger,
	}
}
