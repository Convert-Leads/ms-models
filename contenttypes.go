package models

import (
    "gorm.io/gorm"
)

type ContentType struct {
    gorm.Model
    Name           string `json:"name"`
    Description    string `json:"description"`
}
