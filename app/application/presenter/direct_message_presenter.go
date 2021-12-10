package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type DirectMessagePresenter interface {
	BuildDirectMessageResponse(object model.DirectMessage) (*output.DirectMessageResponse, error)
	BuildDirectMessagesResponse(objects []model.DirectMessage) ([]output.DirectMessageResponse, error)
	BuildDeletedDirectMessageResponse(object model.DirectMessage) (*output.DeletedDirectMessageResponse, error)
}
