package models

type Organisation struct {
	QModel
	Name               string               `json:"n,omitempty"`
	Description        string               `json:"d,omitempty"`
	AppLogo            string               `json:"al,omitempty"`
	AppLogoMedia       *Media               `gorm:"polymorphic:Parent;polymorphicValue:organisations" json:"alm,omitempty"`
	AppName            string               `json:"an,omitempty"`
	WelcomeScreen      *Media               `gorm:"polymorphic:Parent;polymorphicValue:organisations" json:"ws,omitempty"`
	Address            string               `json:"a,omitempty"`
	ContactEmail       string               `json:"ce,omitempty"`
	CountryCode        string               `json:"cc,omitempty"`
	ContactPhone       string               `json:"cp,omitempty"`
	Payments           []StripePayment      `json:"p,omitempty" gorm:"-"`
	SubscriptionLevels []SubscriptionLevel  `json:"sl,omitempty" gorm:"foreignKey:OrganisationId"`
	Users              []User               `json:"u,omitempty" gorm:"foreignKey:OrganisationId"`
	Groups             []Group              `json:"g,omitempty" gorm:"foreignKey:OrganisationId"`
	Content            []Content            `json:"c,omitempty" gorm:"foreignKey:OrganisationId"`
	Tags               []MetadataTag        `json:"t,omitempty" gorm:"foreignKey:OrganisationId"`
	HideCounts         bool                 `json:"hc,omitempty"`
	APIKeys            []OrganisationAPIKey `json:"ak,omitempty" gorm:"foreignKey:OrganisationID"`
	Branding           *Branding            `json:"b,omitempty" gorm:"foreignKey:OrganisationID"`
	OrgSubscriptions   []OrgSubscription    `json:"os,omitempty" gorm:"foreignKey:OrganisationId"`
}
