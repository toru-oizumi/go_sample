package controller

import (
	"go_sample/app/application/interactor"
	"go_sample/app/domain/repository"
	"go_sample/app/interface/presenter_impl"
)

type Controller struct {
	connection repository.Connection
}

func NewController(
	connection repository.Connection,
) *Controller {
	return &Controller{
		connection: connection,
	}
}

func (c *Controller) User() *UserController {
	return &UserController{
		Usecase: &interactor.UserInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewUserPresenter(),
		},
	}
}

func (c *Controller) Group() *GroupController {
	return &GroupController{
		Usecase: &interactor.GroupInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewGroupPresenter(),
		},
	}
}

func (c *Controller) Play() *PlayController {
	return &PlayController{
		Usecase: &interactor.PlayInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewPlayPresenter(),
		},
	}
}

func (c *Controller) Chat() *ChatController {
	return &ChatController{
		Usecase: &interactor.ChatInteractor{
			Connection: c.connection,
			Presenter:  presenter_impl.NewChatPresenter(),
		},
	}
}
