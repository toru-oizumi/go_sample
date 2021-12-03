package repository

import (
	"go_sample/app/domain/model"
)

type UserQuery interface {
	Exists(id model.UserID) (bool, error)
	FindByID(id model.UserID) (*model.User, error)
	List(filter UserFilter) ([]model.User, error)
}

type UserCommand interface {
	UserQuery
	// Interactor(Usecase)内で使用する場合は、service.CreateUserの使用を推奨
	Store(object model.User) (*model.UserID, error)
	Update(object model.User) (*model.UserID, error)
	UpdateGroupByIDs(ids []model.UserID, group model.Group) error
	Delete(id model.UserID) error
}

type UserFilter struct {
	NameLike string
	GroupID  model.GroupID
}
