package models

import (
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) error {
	// AutoMigrate provided models
	if err := DB.AutoMigrate( // enumerate all models here
		&PollVote{},
		&PollOption{},
		&Poll{},
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
		&DeltaContent{},
		&StripeDetail{},
		&PaymentTransaction{},
		&UserRole{},
		&UserBookmark{},
		&Safe{},
		&Group{},
		&Post{},
	); err != nil {
		return err
	}

	return nil
}
