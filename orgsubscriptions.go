package models

import (
	"time"

	"gorm.io/gorm"
)

type OrgSubscriptionLevel struct {
	gorm.Model
	Name            string                      `json:"name"`
	Limits          []OrgSubscriptionLevelLimit `gorm:"foreignKey:OrgSubscriptionLevelId" json:"limits"`
	MediaID         *uint                       `json:"media_id"`
	Media           *Media                      `gorm:"polymorphic:Parent;polymorphicValue:subscriptionlevels;foreignKey:MediaID" json:"media"`
	Description     string                      `json:"description" gorm:"type:text"`
	PriceOptions    []OrgPriceOption            `json:"price_options" gorm:"foreignKey:SubscriptionLevelId"`
	Subscribers     []OrgSubscription           `json:"subscribers" gorm:"foreignKey:SubscriptionLevelId"`
	SubscriberCount int                         `gorm:"-" json:"subscriber_count"`
	Active          bool                        `json:"active" gorm:"default:true"`
	StripeProductID string                      `json:"stripe_product_id"`
}

type OrgPriceOption struct {
	gorm.Model
	SubscriptionLevelId uint    `json:"subscription_level_id"`
	Price               float64 `json:"price"`
	Currency            string  `json:"currency"`
	Duration            int     `json:"duration"`
	DurationUnit        string  `json:"duration_unit"`
	Active              bool    `json:"active"`
	StripePriceID       string  `json:"stripe_price_id"`
}

type OrgSubscriptionLevelLimit struct {
	QModel
	OrgSubscriptionLevelId uint   `json:"org_subscription_level_id"`
	LimitType              string `json:"limit_type"`
	LimitValue             int    `json:"limit_value"`
	Threshold              int    `json:"threshold"`
}

type OrgSubscription struct {
	QModel
	OrganisationId      uint                  `json:"organisation_id"`
	SubscriptionLevel   OrgSubscriptionLevel  `gorm:"foreignKey:SubscriptionLevelId"`
	SubscriptionLevelId uint                  `json:"subscription_level_id"`
	PriceOptionId       uint                  `json:"price_option_id"`
	PriceOption         OrgPriceOption        `gorm:"foreignKey:PriceOptionId" json:"price_option"`
	StartDate           *time.Time            `json:"start_date"`
	EndDate             *time.Time            `json:"end_date"`
	Status              string                `json:"status"`
	StripeSubscription  OrgStripeSubscription `gorm:"foreignKey:OrgSubscriptionId" json:"stripe_subscription"`
	ReferrerId          *uint                 `json:"referrer_id"`
	Referrer            Referrer              `gorm:"foreignKey:ReferrerId" json:"referrer"`
}

type OrgStripeSubscription struct {
	QModel
	OrgSubscriptionId    uint               `json:"org_subscription_id"`
	StripeSubscriptionId string             `json:"stripe_subscription_id"`
	CurrentPeriodEnd     time.Time          `json:"current_period_end"`
	CurrentPeriodStart   time.Time          `json:"current_period_start"`
	CustomerId           string             `json:"customer_id"`
	Payments             []OrgStripePayment `json:"payments"`
}

type OrgStripePayment struct {
	QModel
	OrgStripeSubscriptionId uint      `json:"org_stripe_subscription_id"`
	Amount                  int       `json:"amount"`
	Currency                string    `json:"currency"`
	Status                  string    `json:"status"`
	Timestamp               time.Time `json:"timestamp"`
}

type Referrer struct {
	QModel
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	Commission    float64 `json:"commission"`
	Terms         string  `json:"terms"`
	Active        bool    `json:"active"`
	PayoutMethod  string  `json:"payout_method"`
	PayoutAccount string  `json:"payout_account"`
	Code          string  `json:"code"`
}

type ReferrerPayment struct {
	QModel
	ReferrerId uint    `json:"referrer_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Status     string  `json:"status"`
	Method     string  `json:"method"`
	Account    string  `json:"account"`
}
