package repository

import (
	"go_sample/app/domain/model"
)

type RoomQuery interface {
	FindByID(id model.RoomID) (*model.Room, error)
	List(filter RoomFilter) (model.Rooms, error)
}

type RoomCommand interface {
	RoomQuery
	Store(object model.Room) (*model.Room, error)
	Update(object model.Room) (*model.Room, error)
	DeleteByID(id model.RoomID) error
}

type RoomFilter struct {
	RoomID model.RoomID
}
