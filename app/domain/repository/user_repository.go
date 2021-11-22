package repository

import (
	"go_sample/app/domain/model"
)

type UserQuery interface {
	FindByID(id model.UserID) (*model.User, error)
	List(filter UserFilter) ([]model.User, error)
}

type UserCommand interface {
	UserQuery
	Store(object model.User) (*model.User, error)
	Update(object model.User) (*model.User, error)
	Delete(id model.UserID) error
}

type UserFilter struct {
	NameLike string
	GroupID  model.GroupID
}
