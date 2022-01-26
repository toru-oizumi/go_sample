package repository

import (
	"go_sample/app/domain/model"
)

type AccountQuery interface {
	FindByID(email model.UserID) (*model.Account, error)
	FindByEmail(email model.Email) (*model.Account, error)
	Authenticate(email model.Email, password model.Password) error
}

type AccountCommand interface {
	AccountQuery
	// Interactor(Usecase)内で使用する場合は、userService.Createの使用を推奨
	Store(object model.Account) (*model.UserID, error)
	Activate(email model.Email, currentPassword model.Password, newPassword model.Password) error
	Update(object model.Account) (*model.UserID, error)
	Enable(id model.UserID) error
	Disable(id model.UserID) error
	Delete(id model.UserID) error
}
