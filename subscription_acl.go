package models

import "gorm.io/gorm"

type SubscriptionACL struct {
    gorm.Model
    SubscriptionLevelId uint   `json:"subscription_level_id"`
    Name                string `json:"name"`
    Description         string `json:"description"`
    Filter              string `json:"filter"`
    FilterType          string `json:"filter_type"`
    // If you need to navigate back to SubscriptionLevel from SubscriptionACL, you can include:
    // SubscriptionLevel   SubscriptionLevel `gorm:"foreignKey:SubscriptionLevelId" json:"subscription_level,omitempty"` // Reverse association (optional)
}