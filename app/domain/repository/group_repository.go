package repository

import (
	"go_sample/app/domain/model"
)

type GroupQuery interface {
	FindByID(id model.GroupID) (*model.Group, error)
	List(filter GroupFilter) ([]model.Group, error)
}

type GroupCommand interface {
	GroupQuery
	Store(object model.Group) (*model.Group, error)
	Update(object model.Group) (*model.Group, error)
	Delete(id model.GroupID) error
}

type GroupFilter struct {
	NameLike string
}
