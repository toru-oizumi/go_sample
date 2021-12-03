package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type PlayRDBRecord struct {
	ID            string `gorm:"type:varchar(255);primarykey"`
	Name          string `gorm:"type:varchar(255);unique;not null"`
	OwnerUserID   string
	OwnerUser     UserRDBRecord `gorm:"foreignKey:OwnerUserID"`
	VisitorUserID string
	VisitorUser   UserRDBRecord `gorm:"foreignKey:VisitorUserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"` // gormのデフォルトに則って設定しているが、基本物理削除するので使わない想定
}

func (PlayRDBRecord) TableName() string {
	return "plays"
}

func (r *PlayRDBRecord) ToDomain() (*model.Play, error) {
	group := model.Play{
		ID:            model.PlayID(r.ID),
		Name:          model.PlayName(r.Name),
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

func (r *PlayRDBRecord) FromDomain(d model.Play) PlayRDBRecord {
	return PlayRDBRecord{
		ID:            string(d.ID),
		Name:          string(d.Name),
		OwnerUserID:   string(d.OwnerUserID),
		VisitorUserID: string(d.VisitorUserID),
		CreatedAt:     d.CreatedAt,
		UpdatedAt:     d.UpdatedAt,
	}
}
