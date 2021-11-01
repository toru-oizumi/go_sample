package handler

import (
	"encoding/json"
	"fmt"
	"go_sample/app/domain/model"

	"go_sample/app/interface/controller/logger"
	enum_connection "go_sample/app/interface/controller/ws/enum/connection"

	"github.com/labstack/echo/v4"

	"github.com/gorilla/websocket"
)

type ChatWsHandler struct {
	Logger logger.Logger
}

type ChatRequest struct {
	Message string         `json:"message"`
	SendTo  []model.UserID `json:"sendTo"`
}

func (handler *ChatWsHandler) Handle(c echo.Context) error {
	room_id := c.Param("id")
	// user_idはCognito（というかJWT）から取得する想定
	user_id := model.UserID(c.QueryParam("user_id"))

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	if err := connectionPool.AddConnection(user_id, enum_connection.Chat, room_id, ws); err != nil {
		return err
	}

	defer connectionPool.RemoveConnection(user_id, enum_connection.Chat)

	for {
		// Read
		messageType, p, err := ws.ReadMessage()
		if messageType != websocket.TextMessage {
			return err
		}
		if err != nil {
			return err
		}

		request := new(ChatRequest)
		json.Unmarshal(p, &request)

		// Write
		v, _ := json.Marshal(request)

		fmt.Println(request.Message)
		fmt.Println(request.SendTo)
		fmt.Println(request.SendTo[0])
		fmt.Println(request.SendTo[1])

		connectins := connectionPool.FilterConnectionsByObjective(enum_connection.Chat, room_id)
		if err := sendMessageToConnections(connectins, v); err != nil {
			return err
		}
	}
}
