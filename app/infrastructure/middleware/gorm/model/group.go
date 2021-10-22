package model

import (
	"go_sample/app/domain/model"

	"gorm.io/gorm"
)

type GroupRDBRecord struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);unique;not null"`
}

func (GroupRDBRecord) TableName() string {
	return "groups"
}

func (g *GroupRDBRecord) ToDomain() (*model.Group, error) {
	group := model.Group{
		Id:        model.GroupId(g.ID),
		Name:      model.GroupName(g.Name),
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}

	err := group.Validate()
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRDBRecord) FromDomain(d model.Group) GroupRDBRecord {
	return GroupRDBRecord{
		Model: gorm.Model{
			ID:        uint(d.Id),
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Name: string(d.Name),
	}
}
