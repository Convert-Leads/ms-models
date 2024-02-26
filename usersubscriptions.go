package models

import (
    "gorm.io/gorm"
    "time"
)

type UserSubscription struct {
    gorm.Model
    UserId               uint        `json:"user_id"`
    SubscriptionLevel    SubscriptionLevel `gorm:"foreignKey:SubscriptionLevelId"` 
    SubscriptionLevelId  uint        `json:"subscription_level_id"`
    StartDate            *time.Time `json:"start_date"`
    EndDate              *time.Time `json:"end_date"`
}
