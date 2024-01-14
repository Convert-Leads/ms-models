package models

import (
    "gorm.io/gorm"
)

type UserInteractionType struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}
