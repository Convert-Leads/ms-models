package models

import "gorm.io/gorm"

type SubscriptionACL struct {
    gorm.Model
    SubscriptionLevelID uint   `json:"subscription_level_id"`
    Name               string `json:"name"`
    Description        string `json:"description"`
    Filter             string `json:"filter"`
    FilterType         string `json:"filter_type"`
}
