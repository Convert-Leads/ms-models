package models

import (
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    // AutoMigrate provided models
    if err := db.AutoMigrate(// enumerate all models here
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
