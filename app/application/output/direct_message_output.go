package output

import (
	"go_sample/app/domain/model"
	"time"
)

type DirectMessageFromUserResponse struct {
	ID   model.UserID   `json:"userID" validate:"required"`
	Name model.UserName `json:"name" validate:"required"`
}

type DirectMessageToUserResponse struct {
	ID   model.UserID   `json:"userID" validate:"required"`
	Name model.UserName `json:"name" validate:"required"`
}

type DirectMessageResponse struct {
	ID        model.DirectMessageID         `json:"directMessageID"`
	FromUser  DirectMessageFromUserResponse `json:"fromUser"`
	ToUser    DirectMessageToUserResponse   `json:"toUser"`
	Body      model.DirectMessageBody       `json:"body"`
	CreatedAt time.Time                     `json:"createdAt"`
	UpdatedAt time.Time                     `json:"updatedAt"`
}

type DeletedDirectMessageResponse struct {
	ID model.DirectMessageID `json:"directMessageID"`
}
