package models

import "time"

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"created_at" swaggertype:"integer"`
	UpdatedAt *time.Time `json:"updated_at" swaggertype:"integer"`
	DeletedAt *time.Time `json:"-" swaggertype:"integer"`
}

type QModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
