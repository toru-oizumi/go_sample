package repository

import (
	"go_sample/app/domain/model"
)

type RoomQuery interface {
	FindById(id model.RoomId) (*model.Room, error)
	List(filter RoomFilter) (model.Rooms, error)
}

type RoomCommand interface {
	RoomQuery
	Store(object model.Room) (*model.Room, error)
	Update(object model.Room) (*model.Room, error)
	DeleteById(id model.RoomId) error
}

type RoomFilter struct {
	RoomID model.RoomId
}
