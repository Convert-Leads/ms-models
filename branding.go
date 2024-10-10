package models

import "gorm.io/gorm"

type Branding struct {
	gorm.Model
	OrganisationID uint   `json:"o" gorm:"uniqueIndex"`
	AppIcon        *Media `json:"ai" gorm:"polymorphic:Parent;polymorphicValue:branding_app_icon"`
	AppName        string `json:"an"`
	AppSplash      *Media `json:"as" gorm:"polymorphic:Parent;polymorphicValue:branding_app_splash"`
	AccentColor    string `json:"ac"`
	TextColor      string `json:"tc"`
}
