package models

type Organisation struct {
	QModel
	Name               string               `json:"n,omitempty"`
	Description        string               `json:"d,omitempty"`
	AppLogo            string               `json:"al,omitempty"`
	AppName            string               `json:"an,omitempty"`
	WelcomeScreen      *Media               `gorm:"polymorphic:Parent;polymorphicValue:organisations" json:"ws,omitempty"`
	Address            string               `json:"a,omitempty"`
	ContactEmail       string               `json:"ce,omitempty"`
	ContactPhone       string               `json:"cp,omitempty"`
	StripeApiKey       string               `json:"-"`
	StripeDetails      []StripeDetail       `json:"sd,omitempty" gorm:"foreignKey:OrganisationId"`
	Payments           []PaymentTransaction `json:"p,omitempty" gorm:"foreignKey:OrganisationId"`
	SubscriptionLevels []SubscriptionLevel  `json:"sl,omitempty" gorm:"foreignKey:OrganisationId"`
	Users              []User               `json:"u,omitempty" gorm:"foreignKey:OrganisationId"`
	Groups             []Group              `json:"g,omitempty" gorm:"foreignKey:OrganisationId"`
	Content            []Content            `json:"c,omitempty" gorm:"foreignKey:OrganisationId"`
	Tags               []MetadataTag        `json:"t,omitempty" gorm:"foreignKey:OrganisationId"`
	HideCounts         bool                 `json:"hc,omitempty"`
	APIKeys            []OrganisationAPIKey `json:"ak,omitempty" gorm:"foreignKey:OrganisationID"`
}
