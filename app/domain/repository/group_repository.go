package repository

import (
	"go_sample/app/domain/model"
)

type GroupQuery interface {
	Exists(id model.GroupID) (bool, error)
	ExistsByName(name model.GroupName) (bool, error)
	FindByID(id model.GroupID) (*model.Group, error)
	FindByName(name model.GroupName) (*model.Group, error)
	List(filter GroupFilter) ([]model.Group, error)
}

type GroupCommand interface {
	GroupQuery
	// Interactor(Usecase)内で使用する場合は、groupService.Createの使用を推奨
	Store(object model.Group) (*model.GroupID, error)
	// Interactor(Usecase)内で使用する場合は、groupService.Updateの使用を推奨
	Update(object model.Group) (*model.GroupID, error)
	IncreaseNumberOfMembers(id model.GroupID, num uint) error
	DecreaseNumberOfMembers(id model.GroupID, num uint) error
	// Interactor(Usecase)内で使用する場合は、groupService.Deleteの使用を推奨
	Delete(id model.GroupID) error
}

type GroupFilter struct {
	NameLike string
}
