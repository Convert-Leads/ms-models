package models

import (
    "gorm.io/gorm"
)

type UserRole struct {
    gorm.Model
    UserID int `json:"user_id"`
    RoleID int `json:"role_id"`
}
