package models

import (
    "gorm.io/gorm"
)

type UserRoles struct {
    gorm.Model
    UserID int `json:"user_id"`
    RoleID int `json:"role_id"`
}
