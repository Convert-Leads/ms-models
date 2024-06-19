package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Members     []User `gorm:"many2many:user_groups;" json:"members"`
	Posts       []Post `gorm:"foreignKey:GroupId" json:"posts"`
}
