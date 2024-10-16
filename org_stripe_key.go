package models

type OrgStripeKey struct {
	QModel
	OrganisationID uint   `json:"organisation_id"`
	APIKey         string `json:"api_key" gorm:"type:text"`
	IsLive         bool   `json:"is_live" gorm:"default:false"`
	Active         bool   `json:"active" gorm:"default:false"`
}

type OrgStripeWebhook struct {
	QModel
	OrganisationID uint   `json:"organisation_id"`
	Webhook        string `json:"webhook" gorm:"type:text"`
	WebhookSecret  string `json:"webhook_secret" gorm:"type:text"`
	IsLive         bool   `json:"is_live" gorm:"default:false"`
	Active         bool   `json:"active" gorm:"default:false"`
}

