package models

import (
    "gorm.io/gorm"
)

type Media struct {
    gorm.Model
    OrganisationId uint    `json:"organisation_id"`
    Uri            string `json:"uri"`
}
