package models

import (
    "gorm.io/gorm"
)

type UserRole struct {
    gorm.Model
    UserId uint `json:"user_id"`
    RoleId uint `json:"role_id"`
}
