package handler

import (
	"go_sample/app/application/interactor"
	"go_sample/app/domain/repository"
	"go_sample/app/interface/presenter_impl"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{}
)

type WsControllerr struct {
	connection repository.Connection
}

func NewWsHandler(
	connection repository.Connection,
) *WsControllerr {
	return &WsControllerr{
		connection: connection,
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
