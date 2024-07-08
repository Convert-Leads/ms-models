package models

import (
    "gorm.io/gorm"
)

type ContentType struct {
    QModel
    Name           string `json:"name"`
    Description    string `json:"description"`
}
