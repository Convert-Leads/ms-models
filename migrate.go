package models

import (
    "gorm.io/gorm"
)

func Migrate(DB *gorm.DB) error {
    // AutoMigrate provided models
    if err := DB.AutoMigrate(// enumerate all models here
        &ContentType{},
        &Organisation{},
        &User{},
        &Content{},
        &UserInteraction{},
        &Media{},
        &UserInteractionType{},
        &ContentMetadataTag{},
        &MetadataTag{},
        &Role{},
        &SubscriptionLevel{},
        &UserSubscription{},
        &UserBookmark{},
        &DeltaContent{},
        &StripeDetail{},
        &PaymentTransaction{},
        &SubscriptionACL{},
        &UserRole{},
        &UserBookmark{},
        &Safe{},
        &PriceOption{},
        
    ); err != nil {
        return err
    }

    return nil 
}
