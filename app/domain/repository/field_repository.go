package repository

import (
	"go_sample/app/domain/model"
)

type FieldQuery interface {
	Exists(id model.FieldID) (bool, error)
	FindByID(id model.FieldID) (*model.Field, error)
	List(filter FieldFilter) ([]model.Field, error)
}

type FieldCommand interface {
	FieldQuery
	Store(object model.Field) (*model.FieldID, error)
	Update(object model.Field) (*model.FieldID, error)
	Delete(id model.FieldID) error
}

type FieldFilter struct {
	FieldID model.FieldID
}
