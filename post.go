package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ContentID       uint       `json:"content_id"`
	CreatedByUserId uint       `json:"created_by_user_id"`
	GroupId         uint       `json:"group_id"`
	PostedAt        *time.Time `json:"posted_at"`
	User            User       `gorm:"foreignKey:CreatedByUserId" json:"user"`
	Group           Groups     `gorm:"foreignKey:GroupId" json:"group"`
}
