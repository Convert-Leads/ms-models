package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primary_key" json:"ID"`
	CreatedAt *time.Time     `json:"created_at" swaggertype:"integer"`
	UpdatedAt *time.Time     `json:"updated_at" swaggertype:"integer"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" swaggertype:"integer"`
}

type QModel struct {
	ID        uint           `gorm:"primary_key" json:"ID"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
