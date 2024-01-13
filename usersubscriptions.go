package models

import (
    "gorm.io/gorm"
    "time"
)

type UserSubscriptions struct {
    gorm.Model
    UserID               int        `json:"user_id"`
    SubscriptionLevelID int        `json:"subscription_level_id"`
    StartDate            *time.Time `json:"start_date"`
    EndDate              *time.Time `json:"end_date"`
}
