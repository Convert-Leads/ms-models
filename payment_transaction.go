package models

import (
    "gorm.io/gorm"
)

// PaymentTransaction is a model for a payment transaction
type PaymentTransaction struct {
	gorm.Model
	StripeDetailID  uint       `json:"stripe_detail_id"` // Foreign key to StripeDetail
	OrganisationId  uint       `json:"organisation_id"` // Helps with direct queries for an organisation
	TransactionID   string     `json:"transaction_id"` // Stripe's charge or transaction ID
	Amount          int64      `json:"amount"` // Amount in smallest currency unit (e.g., cents)
	Currency        string     `json:"currency"`
	Status          string     `json:"status"` // e.g., succeeded, failed, pending
	Description     string     `json:"description"` // Optional: Description of the transaction
	PaymentMethodID string     `json:"payment_method_id"` // Stripe PaymentMethod ID
	PaymentIntentID string     `json:"payment_intent_id"` // Stripe PaymentIntent ID
	SubscriptionID  string     `json:"subscription_id"` // Stripe Subscription ID, redundant but useful for direct queries
	StripeDetail    StripeDetail `gorm:"foreignKey:StripeDetailID"`
	Organisation    Organisation `gorm:"foreignKey:OrganisationId"`
}