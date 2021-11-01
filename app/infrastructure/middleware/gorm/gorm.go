package gorm

import (
	"go_sample/app/domain/repository"
	"go_sample/app/infrastructure/middleware/gorm/repository_impl"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) repository.Repository {
	return &dbRepository{
		db: db,
	}
}

type dbRepository struct {
	db *gorm.DB
}

func (r *dbRepository) NewConnection() (repository.Connection, error) {
	return &dbConnection{
		db: r.db,
	}, nil
}

func (r *dbRepository) MustConnection() repository.Connection {
	con, err := r.NewConnection()
	if err != nil {
		panic(err)
	}

	return con
}

type dbConnection struct {
	db *gorm.DB
}

func (con *dbConnection) Close() error {
	// We don't need to close *gorm.DB. No need to do anything.
	return nil
}

func (con *dbConnection) RunTransaction(f func(repository.Transaction) (interface{}, error)) (interface{}, error) {
	tx := con.db.Begin()

	data, err := f(&dbTransaction{db: tx})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (con *dbConnection) User() repository.UserQuery {
	return &repository_impl.UserRepository{Db: con.db}
}
func (con *dbConnection) Group() repository.GroupQuery {
	return &repository_impl.GroupRepository{Db: con.db}
}
func (con *dbConnection) Play() repository.PlayQuery {
	return &repository_impl.PlayRepository{Db: con.db}
}

type dbTransaction struct {
	db *gorm.DB
}

func (tx *dbTransaction) User() repository.UserCommand {
	return &repository_impl.UserRepository{Db: tx.db}
}
func (tx *dbTransaction) Group() repository.GroupCommand {
	return &repository_impl.GroupRepository{Db: tx.db}
}
func (tx *dbTransaction) Play() repository.PlayCommand {
	return &repository_impl.PlayRepository{Db: tx.db}
}
