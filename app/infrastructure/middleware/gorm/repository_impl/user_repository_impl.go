package repository_impl

import (
	"errors"
	"fmt"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	db_model "go_sample/app/infrastructure/middleware/gorm/model"
	"go_sample/app/interface/gateway/db"
	"go_sample/app/utility"
	util_error "go_sample/app/utility/error"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	DB      *gorm.DB
	Service db.DBService
}

func (repo *UserRepository) FindByID(id model.UserID) (*model.User, error) {
	var db_user db_model.UserRDBRecord

	if err := repo.DB.Joins("Group").Take(
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

func (repo *UserRepository) List(filter repository.UserFilter) ([]model.User, error) {
	var db_users []db_model.UserRDBRecord
	var users []model.User

	query := repo.DB.Joins("Group")

	if len(filter.GroupID) != 0 {
		query = query.Where("`group_id` = ?", filter.GroupID)
	}

	if len(filter.NameLike) != 0 {
		query = query.Where("`users`.`name` LIKE ?", fmt.Sprintf("%%%s%%", filter.NameLike))
	}

	if err := query.Find(&db_users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.User{}, nil
		} else {
			return nil, err
		}
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
	db_user.ID = utility.GetUlid()

	if err := repo.DB.Create(&db_user).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	if user, err := repo.FindByID(model.UserID(db_user.ID)); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (repo *UserRepository) Update(object model.User) (*model.User, error) {
	var db_user db_model.UserRDBRecord

	if err := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_user, "`id` = ?", string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}
	db_user.Name = string(object.Name)
	db_user.Age = uint(object.Age)
	db_user.GroupID = string(object.Group.ID)

	if err := repo.DB.Save(&db_user).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	if user, err := repo.FindByID(model.UserID(db_user.ID)); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (repo *UserRepository) Delete(id model.UserID) error {
	var db_user db_model.UserRDBRecord
	if err := repo.DB.Take(&db_user, string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util_error.NewErrRecordNotFound()
		}
		return err
	}

	if err := repo.DB.Delete(&db_user).Error; err != nil {
		return err
	}

	return nil
}
