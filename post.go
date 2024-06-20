package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ContentID        uint              `json:"content_id"`
	CreatedByUserId  uint              `json:"created_by_user_id"`
	GroupId          uint              `json:"group_id"`
	User             User              `gorm:"foreignKey:CreatedByUserId" json:"user"`
	Group            Group             `gorm:"foreignKey:GroupId" json:"group"`
	UserInteractions []UserInteraction `gorm:"polymorphic:Parent;polymorphicValue:posts" json:"user_interactions"`
}
