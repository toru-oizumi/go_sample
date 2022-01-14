package controller

import (
	"encoding/json"
	"go_sample/app/application/usecase"
	"go_sample/app/domain/model"

	"go_sample/app/interface/controller/ws/enum/process"
	"go_sample/app/interface/controller/ws/enum/resource"
	util_error "go_sample/app/utility/error"
)

type FieldWsControllerr struct {
	Usecase usecase.FieldUsecase
}

type FieldWsRequest struct {
	Resource resource.Resource `json:"resource"`
	Process  process.Process   `json:"process"`
}

type FieldWsResponse struct {
	Resource resource.Resource `json:"resource"`
	Process  process.Process   `json:"process"`
	Message  interface{}       `json:"fieldMessage"`
}

func (ctrl *FieldWsControllerr) Handle(user_id model.UserID, message []byte) ([]model.UserID, []byte, error) {
	request := new(FieldWsRequest)
	json.Unmarshal(message, &request)

	res := FieldWsResponse{
		Resource: request.Resource,
		Process:  request.Process,
	}

	switch request.Process {
	case process.Add:
		result, _ := json.Marshal(res)
		return []model.UserID{}, result, nil
	case process.Modify:
		result, _ := json.Marshal(res)
		return []model.UserID{}, result, nil
	case process.Delete:
		result, _ := json.Marshal(res)
		return []model.UserID{}, result, nil
	default:
		// TODO: エラーの返し方
		return nil, nil, util_error.NewErrBadRequest("")
	}
}
