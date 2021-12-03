package handler

import (
	"go_sample/app/domain/repository"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

type WsHandler struct {
	connection repository.Connection
}

func NewWsHandler(
	connection repository.Connection,
) *WsHandler {
	return &WsHandler{
		connection: connection,
	}
}

func (h *WsHandler) Chat() *ChatWsHandler {
	return &ChatWsHandler{}
}

func (h *WsHandler) Field() *FieldWsHandler {
	return &FieldWsHandler{}
}
