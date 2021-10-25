package repository_impl

import (
	"errors"
	"fmt"
	"go_sample/app/domain/model"
	util_error "go_sample/app/utility/error"

	"go_sample/app/domain/repository"
	"go_sample/app/utility"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	db_model "go_sample/app/infrastructure/middleware/gorm/model"
)

type UserRepository struct {
	Db *gorm.DB
}

func (repo *UserRepository) FindById(id model.UserId) (*model.User, error) {
	var db_user db_model.UserRDBRecord

	if err := repo.Db.Joins("Group").Take(
		&db_user,
		fmt.Sprintf("`%s`.`id` = ?", db_user.TableName()),
		string(id),
	).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	if user, err := db_user.ToDomain(); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (repo *UserRepository) List(filter repository.UserFilter) (model.Users, error) {
	var db_users []db_model.UserRDBRecord
	var users model.Users

	if err := repo.Db.Joins("Group").Find(&db_users).Error; err != nil {
		return nil, err
	} else {
		for _, v := range db_users {
			if user, err := v.ToDomain(); err != nil {
				return nil, err
			} else {
				users = append(users, *user)
			}
		}
		return users, nil
	}
}

func (repo *UserRepository) Store(object model.User) (*model.User, error) {
	var db_user db_model.UserRDBRecord
	db_user = db_user.FromDomain(object)
	db_user.Id = utility.GetUlid()

	if err := repo.Db.Create(&db_user).Error; err != nil {
		// ここではGormに依存はしても、DBの種類に依存したくはないが、妥協
		// DBがMySQLの場合
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, util_error.NewErrRecordDuplicate()
		}
		return nil, err
	}

	if user, err := repo.FindById(model.UserId(db_user.Id)); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (repo *UserRepository) Update(object model.User) (*model.User, error) {
	var db_user db_model.UserRDBRecord

	if err := repo.Db.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_user, string(object.Id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}
	db_user.Name = string(object.Name)
	db_user.Age = uint(object.Age)
	db_user.GroupId = string(object.Group.Id)

	if err := repo.Db.Save(&db_user).Error; err != nil {
		// ここではGormに依存はしても、DBの種類に依存したくはないが、妥協
		// DBがMySQLの場合
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, util_error.NewErrRecordDuplicate()
		}
		return nil, err
	}

	if user, err := repo.FindById(model.UserId(db_user.Id)); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (repo *UserRepository) DeleteById(id model.UserId) error {
	var db_user db_model.UserRDBRecord
	if err := repo.Db.Take(&db_user, string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util_error.NewErrRecordNotFound()
		}
		return err
	}

	if err := repo.Db.Delete(&db_user).Error; err != nil {
		return err
	}

	return nil
}
