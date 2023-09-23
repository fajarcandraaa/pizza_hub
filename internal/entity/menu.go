package entity

import "time"

type Menu struct {
	ID        string    `json:"id" gorm:"size:36;not null;unique index;primaryKey"`
	MenuName  string    `json:"name" gorm:"size:50;"`
	MenuCode  string    `json:"code" gorm:"size:20;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
