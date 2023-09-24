package entity

import "time"

type Order struct {
	ID            string    `json:"id" gorm:"size:36;not null;unique index;primaryKey"`
	MenuCode      string    `json:"menu_code"`
	CheffInitials string    `json:"initials" gorm:"size:10;"`
	InProgress    bool      `json:"in_progress" gorm:"default:true"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
