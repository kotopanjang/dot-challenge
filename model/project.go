package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// Project is an entity for table: projects
type Project struct {
	gorm.Model
	ID uint `gorm:"primarykey"`

	Name       string
	ClientName string
	Budget     int
	Progress   float64
	Members    []Member `gorm:"foreignKey:ProjectId"`

	CreatedAt time.Time    `gorm:"->"`
	UpdatedAt time.Time    `gorm:"->"`
	DeletedAt sql.NullTime `gorm:"index;->"`
}
