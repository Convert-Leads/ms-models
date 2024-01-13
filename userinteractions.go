package models

import (
    "gorm.io/gorm"
)

type UserInteractions struct {
    gorm.Model
    UserId              int        `json:"user_id"`
    ContentId           int        `json:"content_id"`
    InteractionTypeId   int        `json:"interaction_type_id"`
    InteractionDatetime *time.Time `json:"interaction_datetime"`
    InteractionDetails  string     `json:"interaction_details"`
}
