package models

import (
    "gorm.io/gorm"
    "time"
)

type UserInteraction struct {
    gorm.Model
    UserId              uint       `json:"user_id"`
    ContentId           uint       `json:"content_id"`
    InteractionType     InteractionType `gorm:"foreignKey:InteractionTypeId"`
    InteractionTypeId   uint       `json:"interaction_type_id"`
    InteractionDatetime *time.Time `json:"interaction_datetime"`
    InteractionDetails  string     `json:"interaction_details"`
    // Including reverse association to Content is typically not necessary for JSON serialization, but can be added if needed:
    // Content             Content    `gorm:"foreignKey:ContentId" json:"content,omitempty"` // Reverse association to Content with JSON annotation (optional)
}