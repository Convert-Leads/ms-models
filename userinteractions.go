package models

import (
	"time"

	"gorm.io/gorm"
)

type UserInteraction struct {
	gorm.Model
	UserId              uint
	InteractionTypeId   uint
	InteractionType     UserInteractionType `gorm:"foreignKey:InteractionTypeId" json:"interaction_type"`
	InteractionDatetime *time.Time          `json:"interaction_datetime"`
	InteractionDetails  string              `json:"interaction_details"`
	User                User                `gorm:"foreignKey:UserId" json:"user"`
	ParentType          string              `json:"parent_type"`
	ParentID            uint                `json:"parent_id"`
	ChildInteractions   []UserInteraction   `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"child_interactions"`
}
