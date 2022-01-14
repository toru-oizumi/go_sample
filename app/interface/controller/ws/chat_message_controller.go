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

type ChatWsControllerr struct {
	Usecase usecase.ChatUsecase
}

type ChatWsRequest struct {
	Resource resource.Resource `json:"resource"`
	Process  process.Process   `json:"process"`
	ChatID   model.ChatID      `json:"chatID"`
}

type ChatWsResponse struct {
	Resource resource.Resource `json:"resource"`
	Process  process.Process   `json:"process"`
	Message  interface{}       `json:"chatMessage"`
}

func (ctrl *ChatWsControllerr) Handle(user_id model.UserID, message []byte) ([]model.UserID, []byte, error) {
	request := new(ChatWsRequest)
	json.Unmarshal(message, &request)

	res := ChatWsResponse{
		Resource: request.Resource,
		Process:  request.Process,
	}
	var err error
	switch request.Process {
	case process.Add:
		input := new(input.CreateChatMessageRequest)
		json.Unmarshal(message, &input)
		input.UserID = user_id

		res.Message, err = ctrl.Usecase.CreateMessage(*input)
	case process.Modify:
		input := new(input.UpdateChatMessageRequest)
		json.Unmarshal(message, &input)
		input.UserID = user_id

		res.Message, err = ctrl.Usecase.UpdateMessage(*input)
	case process.Delete:
		input := new(input.DeleteChatMessageRequest)
		json.Unmarshal(message, &input)
		input.UserID = user_id

		res.Message, err = ctrl.Usecase.DeleteMessage(*input)
	default:
		// TODO: エラーの返し方
		return nil, nil, util_error.NewErrBadRequest("")
	}

	if err != nil {
		return nil, nil, err
	}

	sendUserIDs, err := ctrl.Usecase.FindChatMembers(input.FindChatMembersRequest{ChatID: request.ChatID})
	if err != nil {
		// TODO: エラーの返し方
		return nil, nil, err
	}
	result, _ := json.Marshal(res)
	return sendUserIDs, result, nil
}
