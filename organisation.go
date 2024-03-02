package models

import (
    "gorm.io/gorm"
)

type Organisation struct {
	gorm.Model
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	AppLogo        uint              `json:"app_logo"`
	AppName        string            `json:"app_name"`
	WelcomeScreen  uint              `json:"welcome_screen"`
	Address        string            `json:"address"`
	ContactEmail   string            `json:"contact_email"`
	ContactPhone   string            `json:"contact_phone"`
  StripeApiKey   string            `json:"-"`
	StripeDetails  []StripeDetail    `json:"stripe_details" gorm:"foreignKey:OrganisationID"`
	Payments       []PaymentTransaction `json:"payments" gorm:"foreignKey:OrganisationID"`
}
