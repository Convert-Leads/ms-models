package models

import (
    "gorm.io/gorm"
)

// PaymentTransaction is a model for a payment transaction
type PaymentTransaction struct {
	gorm.Model
	StripeDetailID  uint         `json:"sdid"` // Foreign key to StripeDetail
	OrganisationId  uint         `json:"oid"`  // Helps with direct queries for an organisation
	TransactionID   string       `json:"tid"`  // Stripe's charge or transaction ID
	Amount          int64        `json:"amt"`  // Amount in smallest currency unit (e.g., cents)
	Currency        string       `json:"cur"`
	Status          string       `json:"sts"`  // e.g., succeeded, failed, pending
	Description     string       `json:"desc,omitempty"` // Optional: Description of the transaction
	PaymentMethodID string       `json:"pmid"` // Stripe PaymentMethod ID
	PaymentIntentID string       `json:"piid"` // Stripe PaymentIntent ID
	SubscriptionID  string       `json:"sid"`  // Stripe Subscription ID, redundant but useful for direct queries
	StripeDetail    StripeDetail `json:"-" gorm:"foreignKey:StripeDetailID"`
	Organisation    Organisation `json:"-" gorm:"foreignKey:OrganisationId"`
}
