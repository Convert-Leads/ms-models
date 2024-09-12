package models

import (
	"gorm.io/gorm"
)

type PriceOption struct {
	gorm.Model
	SubscriptionLevelId uint    `json:"subscription_level_id"`
	Price               float64 `json:"price"`
	Currency            string  `json:"currency"`
	Duration            int     `json:"duration"`
	DurationUnit        string  `json:"duration_unit"`
	Active              bool    `json:"active"`
}
