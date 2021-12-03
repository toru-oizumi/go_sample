package handler

import (
	"encoding/json"
	"go_sample/app/domain/model"

	enum_connection "go_sample/app/interface/controller/ws/enum/connection"
	"go_sample/app/interface/gateway/logger"

	"github.com/labstack/echo/v4"

	"github.com/gorilla/websocket"
)

type FieldWsHandler struct {
	Logger logger.WsApiLogger
}

type FieldRequest struct {
	Aaa string `json:"aaa"`
	Bbb int    `json:"bbb"`
	Ccc bool   `json:"ccc"`
}

func (handler *FieldWsHandler) Handle(c echo.Context) error {
	id := c.Param("id")
	// user_idはCognito（というかJWT）から取得する想定
	user_id := model.UserID(c.QueryParam("user_id"))

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	if err := connectionPool.AddConnection(user_id, enum_connection.Field, id, ws); err != nil {
		return err
	}

	defer connectionPool.RemoveConnection(user_id, enum_connection.Field)

	for {
		// Read
		messageType, p, err := ws.ReadMessage()
		if messageType != websocket.TextMessage {
			return err
		}
		if err != nil {
			return err
		}

		request := new(FieldRequest)
		json.Unmarshal(p, &request)

		// Write
		v, _ := json.Marshal(request)

		connectins := connectionPool.FilterConnectionsByObjective(enum_connection.Field, id)
		if err := sendMessageToConnections(connectins, v); err != nil {
			return err
		}
	}
}
