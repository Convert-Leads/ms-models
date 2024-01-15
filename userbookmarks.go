package models

import (
    "gorm.io/gorm"
)

type UserBookmark struct {
    gorm.Model
    UserID    uint `json:"user_id"`
    ContentID uint `json:"content_id"`
}
