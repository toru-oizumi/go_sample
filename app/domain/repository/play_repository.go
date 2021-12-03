package repository

import (
	"go_sample/app/domain/model"
)

type PlayQuery interface {
	Exists(id model.PlayID) (bool, error)
	FindByID(id model.PlayID) (*model.Play, error)
	List(filter PlayFilter) ([]model.Play, error)
}

type PlayCommand interface {
	PlayQuery
	Store(object model.Play) (*model.PlayID, error)
	Update(object model.Play) (*model.PlayID, error)
	Delete(id model.PlayID) error
}

type PlayFilter struct {
	PlayID model.PlayID
}
