package models

import (
    "gorm.io/gorm"
    "time"
)

type UserInteraction struct {
    gorm.Model
    UserId              uint       `json:"user_id"`
    ContentId           uint       `json:"content_id"`
    InteractionType     UserInteractionType `gorm:"foreignKey:InteractionTypeId"`
    InteractionTypeId   uint       `json:"interaction_type_id"`
    InteractionDatetime *time.Time `json:"interaction_datetime"`
    InteractionDetails  string     `json:"interaction_details"`
    User                *User       `gorm:"foreignKey:UserId" json:"user"` 
}