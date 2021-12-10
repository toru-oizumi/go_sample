package repository

import (
	"go_sample/app/domain/model"
)

type DirectMessageQuery interface {
	Exists(id model.DirectMessageID) (bool, error)
	FindByID(id model.DirectMessageID) (*model.DirectMessage, error)
	List(filter DirectMessageFilter) ([]model.DirectMessage, error)
}

type DirectMessageCommand interface {
	DirectMessageQuery
	Store(object model.DirectMessage) (*model.DirectMessageID, error)
	Update(object model.DirectMessage) (*model.DirectMessageID, error)
	Delete(id model.DirectMessageID) error
	DeleteByFromUserID(user_id model.UserID) error
	DeleteByToUserID(user_id model.UserID) error
}

type DirectMessageFilter struct {
	FromUserID model.UserID
	ToUserID   model.UserID
}
