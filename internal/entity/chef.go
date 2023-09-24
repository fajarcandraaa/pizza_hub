package entity

import "time"

type Chef struct {
	ID            string    `json:"id" gorm:"size:36;not null;unique index;primaryKey"`
	CheffInitials string    `json:"initials" gorm:"size:10;unique;"`
	CheffName     string    `json:"name" gorm:"size:50;"`
	Username      string    `json:"username" gorm:"size:100;"`
	Password      string    `json:"password" gorm:"size:100;"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
