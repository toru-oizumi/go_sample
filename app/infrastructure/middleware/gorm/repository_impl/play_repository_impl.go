package repository_impl

import (
	"errors"
	"go_sample/app/domain/model"
	util_error "go_sample/app/utility/error"

	"go_sample/app/domain/repository"
	"go_sample/app/utility"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	db_model "go_sample/app/infrastructure/middleware/gorm/model"
)

type PlayRepository struct {
	Db *gorm.DB
}

func (repo *PlayRepository) FindByID(id model.PlayID) (*model.Play, error) {
	var db_room db_model.PlayRDBRecord

	if err := repo.Db.Take(&db_room, "`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	if room, err := db_room.ToDomain(); err != nil {
		return nil, err
	} else {
		return room, nil
	}
}

func (repo *PlayRepository) List(filter repository.PlayFilter) (model.Plays, error) {
	var db_rooms []db_model.PlayRDBRecord
	var rooms model.Plays

	if err := repo.Db.Find(&db_rooms).Error; err != nil {
		return nil, err
	} else {
		for _, v := range db_rooms {
			if room, err := v.ToDomain(); err != nil {
				return nil, err
			} else {
				rooms = append(rooms, *room)
			}
		}
		return rooms, nil
	}
}

func (repo *PlayRepository) Store(object model.Play) (*model.Play, error) {
	var db_room db_model.PlayRDBRecord
	db_room = db_room.FromDomain(object)
	db_room.ID = utility.GetUlid()

	if err := repo.Db.Create(&db_room).Error; err != nil {
		// ここではGormに依存はしても、DBの種類に依存したくはないが、妥協
		// DBがMySQLの場合
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, util_error.NewErrRecordDuplicate()
		}
		return nil, err
	}

	if room, err := repo.FindByID(model.PlayID(db_room.ID)); err != nil {
		return nil, err
	} else {
		return room, nil
	}
}

func (repo *PlayRepository) Update(object model.Play) (*model.Play, error) {
	var db_room db_model.PlayRDBRecord
	if err := repo.Db.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_room, "`id` = ?", string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}
	db_room.Name = string(object.Name)

	if err := repo.Db.Save(&db_room).Error; err != nil {
		// ここではGormに依存はしても、DBの種類に依存したくはないが、妥協
		// DBがMySQLの場合
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, util_error.NewErrRecordDuplicate()
		}
		return nil, err
	}

	if room, err := repo.FindByID(model.PlayID(db_room.ID)); err != nil {
		return nil, err
	} else {
		return room, nil
	}
}

func (repo *PlayRepository) DeleteByID(id model.PlayID) error {
	var db_room db_model.PlayRDBRecord
	if err := repo.Db.Take(&db_room, string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util_error.NewErrRecordNotFound()
		}
		return err
	}

	if err := repo.Db.Delete(&db_room).Error; err != nil {
		return err
	}

	return nil
}
