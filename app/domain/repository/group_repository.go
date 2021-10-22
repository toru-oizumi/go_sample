package repository

import (
	"go_sample/app/domain/model"
)

type GroupQuery interface {
	FindById(id model.GroupId) (*model.Group, error)
	List(filter GroupFilter) (model.Groups, error)
}

type GroupCommand interface {
	GroupQuery
	Store(group model.Group) (*model.Group, error)
	Update(group model.Group) (*model.Group, error)
	DeleteById(id model.GroupId) error
}

type GroupFilter struct {
	UserID model.UserId
}