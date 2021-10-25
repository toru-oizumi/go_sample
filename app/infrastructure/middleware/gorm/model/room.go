package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type RoomRDBRecord struct {
	Id            string `gorm:"type:varchar(255);primarykey"`
	Name          string `gorm:"type:varchar(255);unique;not null"`
	OwnerUserId   string
	OwnerUser     UserRDBRecord `gorm:"foreignKey:OwnerUserId"`
	VisitorUserId string
	VisitorUser   UserRDBRecord `gorm:"foreignKey:VisitorUserId"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (RoomRDBRecord) TableName() string {
	return "rooms"
}

func (r *RoomRDBRecord) ToDomain() (*model.Room, error) {
	group := model.Room{
		Id:            model.RoomId(r.Id),
		Name:          model.RoomName(r.Name),
		OwnerUserId:   model.UserId(r.OwnerUserId),
		VisitorUserId: model.UserId(r.VisitorUserId),
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}

	err := group.Validate()
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *RoomRDBRecord) FromDomain(d model.Room) RoomRDBRecord {
	return RoomRDBRecord{
		Id:            string(d.Id),
		Name:          string(d.Name),
		OwnerUserId:   string(d.OwnerUserId),
		VisitorUserId: string(d.VisitorUserId),
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}
