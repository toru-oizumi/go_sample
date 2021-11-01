package repository

import (
	"go_sample/app/domain/model"
)

type PlayQuery interface {
	FindByID(id model.PlayID) (*model.Play, error)
	List(filter PlayFilter) (model.Plays, error)
}

type PlayCommand interface {
	PlayQuery
	Store(object model.Play) (*model.Play, error)
	Update(object model.Play) (*model.Play, error)
	DeleteByID(id model.PlayID) error
}

type PlayFilter struct {
	PlayID model.PlayID
}
