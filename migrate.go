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
		&UserBookmark{},
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
		&Media{},
		&OrgStripeKey{},
		&OrgStripeWebhook{},
		&UserSubscription{},
		&StripeSubscription{},
		&StripePayment{},
	}

	for _, model := range models {
		if err := DB.Debug().AutoMigrate(model); err != nil {
			log.Printf("Failed to migrate %T: %v", model, err)
			return err
		}
	}

	return nil
}
