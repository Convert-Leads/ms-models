package models

import (
    "gorm.io/gorm"
)

type UserBookmark struct {
    gorm.Model
    UserID    int `json:"user_id"`
    ContentID int `json:"content_id"`
}
