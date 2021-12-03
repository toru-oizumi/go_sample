package controller

import (
	"go_sample/app/application/usecase"
	"go_sample/app/interface/gateway/logger"
)

type InitialController struct {
	Usecase usecase.InitialUsecase
	Logger  logger.BatchLogger
}

func (ctrl *InitialController) Initialize() error {
	if err := ctrl.Usecase.DataBaseInitialize(); err != nil {
		return ctrl.Logger.Fatal("Failed to initialize the database.", err)
	}
	return nil
}
