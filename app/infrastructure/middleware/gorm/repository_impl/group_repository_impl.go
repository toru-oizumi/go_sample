package repository_impl

import (
	"errors"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"

	util_error "go_sample/app/utility/error"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	db_model "go_sample/app/infrastructure/middleware/gorm/model"
)

type GroupRepository struct {
	Db *gorm.DB
}

func (repo *GroupRepository) FindById(id model.GroupId) (*model.Group, error) {
	var db_group db_model.GroupRDBRecord

	if err := repo.Db.Take(&db_group, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	if group, err := db_group.ToDomain(); err != nil {
		return nil, err
	} else {
		return group, nil
	}
}

func (repo *GroupRepository) List(filter repository.GroupFilter) (model.Groups, error) {
	var groups model.Groups
	if err := repo.Db.Find(&groups).Error; err != nil {
		return nil, err
	} else {
		return groups, nil
	}
}

func (repo *GroupRepository) Store(g model.Group) (*model.Group, error) {
	var db_group db_model.GroupRDBRecord
	db_group = db_group.FromDomain(g)

	if err := repo.Db.Create(&db_group).Error; err != nil {
		// ここではGormに依存はしても、DBの種類に依存したくはないが、妥協
		// DBがMySQLの場合
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, util_error.NewErrRecordDuplicate()
		default:
			return nil, err
		}
	}

	if group, err := db_group.ToDomain(); err != nil {
		return &model.Group{}, err
	} else {
		return group, nil
	}
}

func (repo *GroupRepository) Update(g model.Group) (*model.Group, error) {
	var db_group db_model.GroupRDBRecord
	if err := repo.Db.Take(&db_group, g.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	db_group.Name = string(g.Name)
	if err := repo.Db.Save(&db_group).Error; err != nil {
		// ここではGormに依存はしても、DBの種類に依存したくはないが、妥協
		// DBがMySQLの場合
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, util_error.NewErrRecordDuplicate()
		default:
			return nil, err
		}
	}

	if group, err := db_group.ToDomain(); err != nil {
		return nil, err
	} else {
		return group, nil
	}
}

func (repo *GroupRepository) DeleteById(id model.GroupId) error {
	var db_group db_model.GroupRDBRecord
	if err := repo.Db.Take(&db_group, uint(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util_error.NewErrRecordNotFound()
		}
		return err
	}

	if err := repo.Db.Delete(&db_group).Error; err != nil {
		return nil
	}

	return nil
}
