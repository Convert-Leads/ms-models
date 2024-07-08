package models

import (
    "gorm.io/gorm"
)

type Media struct {
    models.Qmodel
    OrganisationId uint   `json:"-"`
    Uri            string `json:"uri"`
    ParentType          string              `json:"-"`
    ParentID            uint                `json:"-"`
}