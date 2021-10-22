package repository

import (
	"go_sample/app/domain/model"
)

type UserQuery interface {
	FindById(id model.UserId) (*model.User, error)
	List(filter UserFilter) (model.Users, error)
}

type UserCommand interface {
	UserQuery
	Store(user model.User) (*model.User, error)
	Update(user model.User) (*model.User, error)
	DeleteById(id model.UserId) error
}

type UserFilter struct {
	NameLike string
}
