package models

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) error {
	models := []interface{}{
		&Poll{},
		&PollOption{},
		&PollVote{},
		&ContentType{},
		&Organisation{},
		&User{},
		&CallToAction{},
		&Content{},
		&Collection{},
		&Chapter{},
		&ChapterContent{},
		&UserInteraction{},
		&Media{},
		&UserInteractionType{},
		&ContentMetadataTag{},
		&MetadataTag{},
		&Role{},
		&SubscriptionLevel{},
		&PriceOption{},
		&UserSubscription{},
		&UserBookmark{},
		&UserRole{},
		&Safe{},
		&Group{},
		&Post{},
		&NewsletterTemplate{},
		&Newsletter{},
		&NewsletterElement{},
		&NewsletterElementDecoration{},
		&ContentInteraction{},
		&CollectionInteraction{},
		&OrganisationAPIKey{},
		&Branding{},
		&OrgStripeKey{},
		&OrgStripeWebhook{},
		&StripeSubscription{},
		&StripePayment{},
		&OrgSubscriptionLevel{},
		&OrgPriceOption{},
		&OrgSubscriptionLevelLimit{},
		&OrgSubscription{},
		&OrgStripeSubscription{},
		&OrgStripePayment{},
		&Visit{},
		&SignUpDetails{},
		&ReferrerPayment{},

	for _, model := range models {
		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("Failed to migrate %T: %v", model, err)
			return err
		}
	}

	return nil
}
