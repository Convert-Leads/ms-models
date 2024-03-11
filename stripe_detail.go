package models

import (
	"gorm.io/gorm"
)

// StripeDetail is a model to represent the stripe_detail table.
type StripeDetail struct {
	gorm.Model
	OrganisationId uint                  `json:"organisation_id"`
	CustomerID     string                `json:"customer_id"` // Stripe Customer ID
	SubscriptionID string                `json:"subscription_id"` // Stripe Subscription ID
	Active         string                  `json:"active"` // Indicates if the subscription is active
	// Organisation   Organisation          `gorm:"foreignKey:OrganisationId"`
	Payments       []PaymentTransaction  `json:"payments" gorm:"foreignKey:StripeDetailID"`
}
