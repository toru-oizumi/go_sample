package handler

import (
	"encoding/json"
	"fmt"
	"go_sample/app/application/usecase"
	"go_sample/app/domain/model"

	"github.com/labstack/echo/v4"

	"github.com/gorilla/websocket"
)

type FieldWsControllerr struct {
	Usecase usecase.FieldUsecase
}

type FieldRequest struct {
	Aaa string `json:"aaa"`
	Bbb int    `json:"bbb"`
	Ccc bool   `json:"ccc"`
}

func (handler *FieldWsControllerr) Handle(c echo.Context) error {
	// user_idはCognito（というかJWT）から取得する想定
	user_id := model.UserID(c.QueryParam("userID"))
	// TODO: UserIDの存在確認

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	connectionPool.AddConnection(user_id, ws)

	defer connectionPool.RemoveConnection(user_id)

	for {
		chat_id := model.ChatID(c.Param("id"))
		// TODO: chat_idの存在確認
		// TODO: 送信先 UserIDのリスト作成
		var sendUserIDs []model.UserID
		sendUserIDs = append(sendUserIDs, user_id)

		// Read
		messageType, p, err := ws.ReadMessage()
		if messageType != websocket.TextMessage {
			// TODO: Binaryは受け入れらませんerror
			return err
		}
		if err != nil {
			// TODO: それ以外のエラー
			return err
		}

		request := new(WsChatRequest)
		request.ChatID = chat_id
		json.Unmarshal(p, &request)

		// Write
		v, _ := json.Marshal(request)

		fmt.Println(request.Process)

		connectins := connectionPool.FilterConnectionsByUserIDs(sendUserIDs)
		if err := sendMessageToConnections(connectins, v); err != nil {
			return err
		}
	}
}
