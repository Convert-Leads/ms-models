package models

import (
    "gorm.io/gorm"
)

type SubscriptionLevel struct {
    gorm.Model
    Name           string             `json:"name"`
    Description    string             `json:"description"`
    Price          float64            `json:"price"`
    Currency       string             `json:"currency"`
    Duration       int                `json:"duration"`
    DurationUnit   string             `json:"duration_unit"` 
    OrganisationId uint               `json:"organisation_id"`
    ACLs           []SubscriptionACL  `gorm:"foreignKey:SubscriptionLevelId" json:"acls"` 
}
