package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type RoomRDBRecord struct {
	ID            string `gorm:"type:varchar(255);primarykey"`
	Name          string `gorm:"type:varchar(255);unique;not null"`
	OwnerUserID   string
	OwnerUser     UserRDBRecord `gorm:"foreignKey:OwnerUserID"`
	VisitorUserID string
	VisitorUser   UserRDBRecord `gorm:"foreignKey:VisitorUserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (RoomRDBRecord) TableName() string {
	return "rooms"
}

func (r *RoomRDBRecord) ToDomain() (*model.Room, error) {
	group := model.Room{
		ID:            model.RoomID(r.ID),
		Name:          model.RoomName(r.Name),
		OwnerUserID:   model.UserID(r.OwnerUserID),
		VisitorUserID: model.UserID(r.VisitorUserID),
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
		ID:            string(d.ID),
		Name:          string(d.Name),
		OwnerUserID:   string(d.OwnerUserID),
		VisitorUserID: string(d.VisitorUserID),
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}
