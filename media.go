package models

import (
    "gorm.io/gorm"
)

type Media struct {
    gorm.Model
    OrganisationId uint   `json:"organisation_id"`
    Uri            string `json:"uri"`
    ParentType          string              `json:"parent_type"`
    ParentID            uint                `json:"parent_id"`
}