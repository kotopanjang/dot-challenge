package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// Member is an entity for table: members
type Member struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	Title     string
	ProjectId string

	CreatedAt time.Time    `gorm:"->"`
	UpdatedAt time.Time    `gorm:"->"`
	DeletedAt sql.NullTime `gorm:"index;->"`
}
