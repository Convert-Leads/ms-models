package models

import (
    "gorm.io/gorm"
)

type Roles struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}
