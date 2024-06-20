package models

import (
	"gorm.io/gorm"
)

type Organisation struct {
	gorm.Model
	Name               string               `json:"name"`
	Description        string               `json:"description"`
	AppLogo            uint                 `json:"app_logo"`
	AppName            string               `json:"app_name"`
	WelcomeScreen      uint                 `json:"welcome_screen"`
	Address            string               `json:"address"`
	ContactEmail       string               `json:"contact_email"`
	ContactPhone       string               `json:"contact_phone"`
	StripeApiKey       string               `json:"-"`
	StripeDetails      []StripeDetail       `json:"stripe_details" gorm:"foreignKey:OrganisationId"`
	Payments           []PaymentTransaction `json:"payments" gorm:"foreignKey:OrganisationId"`
	SubscriptionLevels []SubscriptionLevel  `json:"subscription_levels" gorm:"foreignKey:OrganisationId"`
	Users              []User               `json:"users" gorm:"foreignKey:OrganisationId"`
	Groups             []Group              `json:"groups" gorm:"foreignKey:OrganisationId"`
	Content            []Content            `json:"content" gorm:"foreignKey:OrganisationId"`
	Tags               []MetadataTag        `json:"tags" gorm:"foreignKey:OrganisationId"`
}
