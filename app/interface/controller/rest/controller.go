package controller

import (
	"go_sample/app/application/interactor"
	"go_sample/app/domain/repository"
	"go_sample/app/interface/gateway/logger"
	"go_sample/app/interface/presenter_impl"
)

type Controller struct {
	connection repository.Connection
	logger     logger.RestApiLogger
}

func NewController(
	connection repository.Connection,
	logger logger.RestApiLogger,
) *Controller {
	return &Controller{
		connection: connection,
		logger:     logger,
	}
}

func (c *Controller) User() *UserController {
	return &UserController{
		Usecase: &interactor.UserInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewUserPresenter(),
		},
		Logger: c.logger,
	}
}

func (c *Controller) Group() *GroupController {
	return &GroupController{
		Usecase: &interactor.GroupInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewGroupPresenter(),
		},
		Logger: c.logger,
	}
}

func (c *Controller) Play() *PlayController {
	return &PlayController{
		Usecase: &interactor.PlayInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewPlayPresenter(),
		},
		Logger: c.logger,
	}
}

func (c *Controller) Chat() *ChatController {
	return &ChatController{
		Usecase: &interactor.ChatInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewChatPresenter(),
		},
		Logger: c.logger,
	}
}
