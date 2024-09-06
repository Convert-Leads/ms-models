package models

import (
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) error {
	// AutoMigrate provided models
	if err := DB.AutoMigrate( // enumerate all models here
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
		&StripeDetail{},
		&PaymentTransaction{},
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
	); err != nil {
		return err
	}

	return nil
}
