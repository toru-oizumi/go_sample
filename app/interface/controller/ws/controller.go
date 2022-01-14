package controller

import (
	"encoding/json"
	"go_sample/app/application/interactor"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	"go_sample/app/interface/controller/ws/enum/process"
	"go_sample/app/interface/controller/ws/enum/resource"
	"go_sample/app/interface/presenter_impl"
	"net/http"
	"strings"

	util_error "go_sample/app/utility/error"

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

var (
	upgrader = websocket.Upgrader{}
)

type WsControllerr struct {
	connection repository.Connection
}

func NewWsController(
	connection repository.Connection,
) *WsControllerr {
	return &WsControllerr{
		connection: connection,
	}
}

type WsRequest struct {
	Resource resource.Resource `json:"resource"`
	Process  process.Process   `json:"process"`
}

func (ctrl *WsControllerr) Handle(c echo.Context) error {
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
		messageType, message, err := ws.ReadMessage()
		if messageType != websocket.TextMessage {
			// TODO: Binaryは受け入れらませんerror
			return err
		}
		if err != nil {
			// TODO: それ以外のエラー
			return err
		}

		request := new(WsRequest)
		json.Unmarshal(message, &request)

		var ids []model.UserID
		var result []byte
		switch request.Resource {
		case resource.Chat:
			ids, result, err = ctrl.Chat().Handle(user_id, message)
		case resource.DirectMessage:
			ids, result, err = ctrl.DirectMessage().Handle(user_id, message)
		case resource.Field:
			ids, result, err = ctrl.Field().Handle(user_id, message)
		default:
			// TODO: エラーの返し方
			return util_error.NewErrBadRequest("")
		}

		if err != nil {
			return err
		}
		connectins := connectionPool.FilterConnectionsByUserIDs(ids)
		sendMessageToConnections(connectins, result)
	}
}

func (c *WsControllerr) Chat() *ChatWsControllerr {
	return &ChatWsControllerr{
		Usecase: &interactor.ChatInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewChatPresenter(),
		},
	}
}

func (c *WsControllerr) Field() *FieldWsControllerr {
	return &FieldWsControllerr{
		Usecase: &interactor.FieldInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewFieldPresenter(),
		},
	}
}

func (c *WsControllerr) DirectMessage() *DirectMessageWsControllerr {
	return &DirectMessageWsControllerr{
		Usecase: &interactor.DirectMessageInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewDirectMessagePresenter(),
		},
	}
}
