package handler

import (
	"go_sample/app/domain/repository"
	"go_sample/app/interface/controller/logger"
)

type WsHandler struct {
	connection repository.Connection
	logger     logger.Logger
}

func NewWsHandler(
	connection repository.Connection,
	logger logger.Logger,
) *WsHandler {
	return &WsHandler{
		connection: connection,
		logger:     logger,
	}
}

func (h *WsHandler) Room() *RoomWsHandler {
	return &RoomWsHandler{
		Logger: h.logger,
	}
}

// func (c *WsHandler) Room() *UserController {
// 	return &UserController{
// 		Usecase: &interactor.UserInteractor{
// 			Connection: c.connection,
// 			Presenter:  presenter_impl.NewUserPresenter(),
// 		},
// 		Logger: c.logger,
// 	}
// }
