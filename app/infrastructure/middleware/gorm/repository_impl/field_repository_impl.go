package repository_impl

import (
	"errors"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	db_model "go_sample/app/infrastructure/middleware/gorm/model"
	"go_sample/app/interface/gateway/db"
	"go_sample/app/utility"
	util_error "go_sample/app/utility/error"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FieldRepository struct {
	DB      *gorm.DB
	Service db.DBService
}

func (repo *FieldRepository) Exists(id model.FieldID) (bool, error) {
	var db_field db_model.FieldRDBRecord

	if err := repo.DB.Select("`id`").Take(&db_field, "`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *FieldRepository) FindByID(id model.FieldID) (*model.Field, error) {
	var db_field db_model.FieldRDBRecord

	if err := repo.DB.Take(&db_field, "`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	if field, err := db_field.ToDomain(); err != nil {
		return nil, err
	} else {
		return field, nil
	}
}

func (repo *FieldRepository) List(filter repository.FieldFilter) ([]model.Field, error) {
	db_fields := []db_model.FieldRDBRecord{}
	fields := []model.Field{}

	if err := repo.DB.Find(&db_fields).Error; err != nil {
		return nil, err
	} else {
		for _, v := range db_fields {
			if field, err := v.ToDomain(); err != nil {
				return nil, err
			} else {
				fields = append(fields, *field)
			}
		}
		return fields, nil
	}
}

func (repo *FieldRepository) Store(object model.Field) (*model.FieldID, error) {
	var db_field db_model.FieldRDBRecord
	db_field = db_field.FromDomain(object)
	// IDは設定が無ければ生成する
	if len(db_field.ID) <= 0 {
		db_field.ID = utility.GetUlid()
	}

	if err := repo.DB.Create(&db_field).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	field_id := model.FieldID(db_field.ID)
	return &field_id, nil
}

func (repo *FieldRepository) Update(object model.Field) (*model.FieldID, error) {
	var db_field db_model.FieldRDBRecord
	if err := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_field, "`id` = ?", string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}
	db_field.Name = string(object.Name)

	if err := repo.DB.Save(&db_field).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	field_id := model.FieldID(db_field.ID)
	return &field_id, nil
}

func (repo *FieldRepository) Delete(id model.FieldID) error {
	var db_field db_model.FieldRDBRecord
	if err := repo.DB.Unscoped().Delete(&db_field, "`id` = ?", string(id)).Error; err != nil {
		return err
	}

	return nil
}
