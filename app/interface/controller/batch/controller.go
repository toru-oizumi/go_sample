package controller

import (
	"go_sample/app/application/interactor"
	"go_sample/app/domain/repository"
	"go_sample/app/interface/gateway/logger"
)

type Controller struct {
	connection repository.Connection
	logger     logger.BatchLogger
}

func NewController(
	connection repository.Connection,
	logger logger.BatchLogger,
) *Controller {
	return &Controller{
		connection: connection,
		logger:     logger,
	}
}

func (c *Controller) Initial() *InitialController {
	return &InitialController{
		Usecase: &interactor.InitialInteractor{
			Connection: c.connection,
		},
		Logger: c.logger,
	}
}
