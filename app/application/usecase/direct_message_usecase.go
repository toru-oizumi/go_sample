package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type DirectMessageUsecase interface {
	FindMessages(request input.FindDirectMessagesRequest) ([]output.DirectMessageResponse, error)
	CreateMessage(request input.CreateDirectMessageRequest) (*output.DirectMessageResponse, error)
	UpdateMessage(request input.UpdateDirectMessageRequest) (*output.DirectMessageResponse, error)
	DeleteMessage(request input.DeleteDirectMessageRequest) (*output.DeletedDirectMessageResponse, error)
}
