package models

import (
    "gorm.io/gorm"
)

func Migrate(DB *gorm.DB) error {
    // AutoMigrate provided models
    if err := DB.AutoMigrate(// enumerate all models here
        &UserInteraction{},
        &Media{},
        &UserInteractionType{},
        &Organisation{},
        &User{},
        &Content{},
        &ContentType{},
        &ContentMetadataTag{},
        &MetadataTag{},
        &Role{},
        &SubscriptionLevel{},
        &UserSubscription{},
        &UserBookmark{},
        ); err != nil {
        return err
    }

    return nil 
}
