package models

import (
    "gorm.io/gorm"
)

type SubscriptionLevels struct {
    gorm.Model
    Name            string  `json:"name"`
    Description     string  `json:"description"`
    Price           float64 `json:"price"`
    Currency        string  `json:"currency"`
    Duration        int     `json:"duration"`
    Duration_unit   string  `json:"duration_unit"`
    OrganisationID int     `json:"organisation_id"`
}
