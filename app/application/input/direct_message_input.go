package input

import (
	"go_sample/app/domain/model"
)

type FindDirectMessagesRequest struct {
	FromUserID model.UserID `param:"fromUserID" query:"fromUserID" validate:"required"`
	ToUserID   model.UserID `param:"toUserID" query:"toUserID" validate:"required"`
	// TODO: Pagenationが必要になるはず Limit,OffsetでなくID指定による形式かな
}

type CreateDirectMessageRequest struct {
	FromUserID model.UserID            `json:"fromUserID" validate:"required"`
	ToUserID   model.UserID            `json:"toUserID" validate:"required"`
	Message    model.DirectMessageBody `json:"message" validate:"required"`
}

type UpdateDirectMessageRequest struct {
	ID         model.DirectMessageID   `json:"directMessageID" validate:"required"`
	FromUserID model.UserID            `json:"fromUserID" validate:"required"`
	ToUserID   model.UserID            `json:"toUserID" validate:"required"`
	Message    model.DirectMessageBody `json:"message" validate:"required"`
}

type DeleteDirectMessageRequest struct {
	ID         model.DirectMessageID `json:"directMessageID" validate:"required"`
	FromUserID model.UserID          `json:"fromUserID" validate:"required"`
	ToUserID   model.UserID          `json:"toUserID" validate:"required"`
}
