package handler

import (
	"fmt"

	"go_sample/app/interface/controller/logger"
	// "go_sample/app/interface/controller/context"

	"github.com/labstack/echo/v4"

	"github.com/gorilla/websocket"
)

type RoomWsHandler struct {
	// Usecase usecase.GroupUsecase
	Logger logger.Logger
}

var (
	upgrader = websocket.Upgrader{}
)

func (handler *RoomWsHandler) Do(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello!!!!, Client!!!!!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
