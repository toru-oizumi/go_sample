package handler

import (
	"encoding/json"
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/domain/model"
	"net/http"
	"strings"

	enum_process "go_sample/app/interface/controller/ws/enum/process"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

// ログイン時
// ユーザ毎のコネクションを作成する
// ALLチャットと、GroupチャットのメッセージをN件取得 => ログイン時にとりあえず最新化する感じ
// TODO: 差分が出ないようにするにばコネクションと全件取得をどうすれば良いか？

// ログアウト時にコネクションを削除する

// 個別チャットを選ぶと、そのチャットのメッセージをN件取得

// ADDすると
// DBにメッセージを書き込む
// ChatIDから送信先のUserIDリストを取得する
// 取得したUserIDリストからConnectionリストを取得する
// 取得したConnectionリスト宛にメッセージを送信する

type ChatWsControllerr struct {
	Usecase usecase.ChatUsecase
}

type WsChatRequest struct {
	Process enum_process.ChatProcess `json:"process"`
	ChatID  model.ChatID             `json:"chatID"`
}

type WsChatResponse struct {
	Process     enum_process.ChatProcess `json:"process"`
	ChatMessage interface{}              `json:"chatMessage"`
}

func (ctrl *ChatWsControllerr) Handle(c echo.Context) error {
	// user_idはCognito（というかJWT）から取得する想定
	headers := c.Request().Header[http.CanonicalHeaderKey("authorization")]
	dummy_jwt := strings.Replace(headers[0], "Bearer ", "", 1)
	user_id := model.UserID(dummy_jwt)

	// TODO: UserIDの存在確認

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	connectionPool.AddConnection(user_id, ws)

	defer connectionPool.RemoveConnection(user_id)

	for {
		// Read
		messageType, body, err := ws.ReadMessage()
		if messageType != websocket.TextMessage {
			// TODO: Binaryは受け入れらませんerror
			return err
		}
		if err != nil {
			// TODO: それ以外のエラー
			return err
		}

		// TODO: エラーの返し方
		// TODO: バリデーションどうするか？

		request := new(WsChatRequest)
		json.Unmarshal(body, &request)

		var result []byte
		switch request.Process {
		case enum_process.Add:
			chat_req := new(input.CreateChatMessageRequest)
			json.Unmarshal(body, &chat_req)
			chat_req.UserID = user_id
			chat_req.ChatID = request.ChatID
			if message, err := ctrl.Usecase.CreateMessage(*chat_req); err != nil {
				return err
			} else {
				result, _ = json.Marshal(
					WsChatResponse{
						Process:     enum_process.Add,
						ChatMessage: message,
					})
			}
		case enum_process.Modify:
			chat_req := new(input.UpdateChatMessageRequest)
			json.Unmarshal(body, &chat_req)
			chat_req.UserID = user_id
			chat_req.ChatID = request.ChatID
			if message, err := ctrl.Usecase.UpdateMessage(*chat_req); err != nil {
				return err
			} else {
				result, _ = json.Marshal(
					WsChatResponse{
						Process:     enum_process.Modify,
						ChatMessage: message,
					})
			}
		case enum_process.Delete:
			chat_req := new(input.DeleteChatMessageRequest)
			json.Unmarshal(body, &chat_req)
			chat_req.UserID = user_id
			chat_req.ChatID = request.ChatID
			if message, err := ctrl.Usecase.DeleteMessage(*chat_req); err != nil {
				return err
			} else {
				result, _ = json.Marshal(
					WsChatResponse{
						Process:     enum_process.Delete,
						ChatMessage: message,
					})
			}
		default:
			// TODO: エラーの返し方
			return nil
		}

		sendUserIDs, err := ctrl.Usecase.FindChatMembers(input.FindChatMembersRequest{ChatID: request.ChatID})
		if err != nil {
			// TODO: エラーの返し方
			return err
		}

		connectins := connectionPool.FilterConnectionsByUserIDs(sendUserIDs)
		if err := sendMessageToConnections(connectins, result); err != nil {
			return err
		}
	}
}
