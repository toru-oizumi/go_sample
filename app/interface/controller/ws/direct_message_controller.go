package controller

import (
	"encoding/json"
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/domain/model"

	"go_sample/app/interface/controller/ws/enum/process"
	"go_sample/app/interface/controller/ws/enum/resource"
	util_error "go_sample/app/utility/error"
)

type DirectMessageWsControllerr struct {
	Usecase usecase.DirectMessageUsecase
}

type DirectMessageWsRequest struct {
	Resource resource.Resource `json:"resource"`
	Process  process.Process   `json:"process"`
	ToUserID model.UserID      `json:"toUserID"`
}

type DirectMessageWsResponse struct {
	Resource resource.Resource `json:"resource"`
	Process  process.Process   `json:"process"`
	Message  interface{}       `json:"directMessage"`
}

func (ctrl *DirectMessageWsControllerr) Handle(user_id model.UserID, message []byte) ([]model.UserID, []byte, error) {
	request := new(DirectMessageWsRequest)
	json.Unmarshal(message, &request)

	res := DirectMessageWsResponse{
		Resource: request.Resource,
		Process:  request.Process,
	}
	var err error
	switch request.Process {
	case process.Add:
		input := new(input.CreateDirectMessageRequest)
		json.Unmarshal(message, &input)
		input.FromUserID = user_id

		res.Message, err = ctrl.Usecase.CreateMessage(*input)
	case process.Modify:
		input := new(input.UpdateDirectMessageRequest)
		json.Unmarshal(message, &input)
		input.FromUserID = user_id

		res.Message, err = ctrl.Usecase.UpdateMessage(*input)
	case process.Delete:
		input := new(input.DeleteDirectMessageRequest)
		json.Unmarshal(message, &input)
		input.FromUserID = user_id

		res.Message, err = ctrl.Usecase.DeleteMessage(*input)
	default:
		// TODO: エラーの返し方
		return nil, nil, util_error.NewErrBadRequest("")
	}

	if err != nil {
		return nil, nil, err
	}

	ids := []model.UserID{user_id, request.ToUserID}
	result, _ := json.Marshal(res)
	return ids, result, nil
}
