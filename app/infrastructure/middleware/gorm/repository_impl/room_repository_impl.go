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

type RoomRepository struct {
	Db *gorm.DB
}

func (repo *RoomRepository) FindById(id model.RoomId) (*model.Room, error) {
	var db_room db_model.RoomRDBRecord

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

func (repo *RoomRepository) List(filter repository.RoomFilter) (model.Rooms, error) {
	var db_rooms []db_model.RoomRDBRecord
	var rooms model.Rooms

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

func (repo *RoomRepository) Store(object model.Room) (*model.Room, error) {
	var db_room db_model.RoomRDBRecord
	db_room = db_room.FromDomain(object)
	db_room.Id = utility.GetUlid()

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

	if room, err := repo.FindById(model.RoomId(db_room.Id)); err != nil {
		return nil, err
	} else {
		return room, nil
	}
}

func (repo *RoomRepository) Update(object model.Room) (*model.Room, error) {
	var db_room db_model.RoomRDBRecord
	if err := repo.Db.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_room, "`id` = ?", string(object.Id)).Error; err != nil {
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

	if room, err := repo.FindById(model.RoomId(db_room.Id)); err != nil {
		return nil, err
	} else {
		return room, nil
	}
}

func (repo *RoomRepository) DeleteById(id model.RoomId) error {
	var db_room db_model.RoomRDBRecord
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
