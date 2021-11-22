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

type GroupRepository struct {
	DB      *gorm.DB
	Service db.DBService
}

func (repo *GroupRepository) FindByID(id model.GroupID) (*model.Group, error) {
	var db_group db_model.GroupRDBRecord

	if err := repo.DB.Take(&db_group, "`id` = ?", string(id)).Error; err != nil {
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

func (repo *GroupRepository) List(filter repository.GroupFilter) ([]model.Group, error) {
	var db_groups []db_model.GroupRDBRecord
	var groups []model.Group

	query := repo.DB

	if len(filter.NameLike) != 0 {
		query = query.Where("`groups`.`name` LIKE ?", fmt.Sprintf("%%%s%%", filter.NameLike))
	}

	if err := query.Find(&db_groups).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Group{}, nil
		} else {
			return nil, err
		}
	} else {
		for _, v := range db_groups {
			if group, err := v.ToDomain(); err != nil {
				return nil, err
			} else {
				groups = append(groups, *group)
			}
		}
		return groups, nil
	}
}

func (repo *GroupRepository) Store(object model.Group) (*model.Group, error) {
	var db_group db_model.GroupRDBRecord
	db_group = db_group.FromDomain(object)
	db_group.ID = utility.GetUlid()

	if err := repo.DB.Create(&db_group).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	if group, err := db_group.ToDomain(); err != nil {
		return &model.Group{}, err
	} else {
		return group, nil
	}
}

func (repo *GroupRepository) Update(object model.Group) (*model.Group, error) {
	var db_group db_model.GroupRDBRecord
	if err := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_group, "`id` = ?", string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}
	db_group.Name = string(object.Name)

	if err := repo.DB.Save(&db_group).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	if group, err := db_group.ToDomain(); err != nil {
		return nil, err
	} else {
		return group, nil
	}
}

func (repo *GroupRepository) Delete(id model.GroupID) error {
	var db_group db_model.GroupRDBRecord
	if err := repo.DB.Take(&db_group, string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util_error.NewErrRecordNotFound()
		}
		return err
	}

	if err := repo.DB.Delete(&db_group).Error; err != nil {
		return nil
	}

	return nil
}
