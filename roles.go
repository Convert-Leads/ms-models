package models

import (
    "gorm.io/gorm"
)

type Role struct {
    QModel
    Name        string `json:"n"`
    Description string `json:"d"`
}
