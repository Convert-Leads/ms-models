package models

import (
    "time"
)

type UserSubscription struct {
    QModel
    UserId               uint        `json:"user_id"`
    SubscriptionLevel    SubscriptionLevel `gorm:"foreignKey:SubscriptionLevelId"` 
    SubscriptionLevelId  uint        `json:"subscription_level_id"`
    PriceOptionId        uint        `json:"price_option_id"`
    PriceOption          PriceOption `gorm:"foreignKey:PriceOptionId" json:"price_option"`
    StartDate            *time.Time `json:"start_date"`
    EndDate              *time.Time `json:"end_date"`
    Status               string      `json:"status"`
    StripeSubscription StripeSubscription `gorm:"foreignKey:SubscriptionId" json:"stripe_subscription"`
}

type StripeSubscription struct {
    QModel
    SubscriptionId     string          `json:"subscription_id"`
    Status             string          `json:"status"`
    CurrentPeriodEnd   time.Time       `json:"current_period_end"`
    CurrentPeriodStart time.Time       `json:"current_period_start"`
    CustomerId         string          `json:"customer_id"`
    Payments           []StripePayment `gorm:"foreignKey:SubscriptionId;references:SubscriptionId" json:"payments"`
}

type StripePayment struct {
    QModel
    SubscriptionId string    `json:"subscription_id"`
    Amount         int       `json:"amount"`
    Currency       string    `json:"currency"`
    Status         string    `json:"status"`
    Timestamp      time.Time `json:"timestamp"`
}
