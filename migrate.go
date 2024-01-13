package models

import (
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    // AutoMigrate provided models
    if err := db.AutoMigrate(// enumerate all models here
        &UserInteractions{},
        &Media{},
        &UserInteractionTypes{},
        &Organisations{},
        &Users{},
        &Content{},
        &ContentTypes{},
        &ContentMetadataTags{},
        &MetadataTags{},
        &Roles{},
        &UserRoles{},
        &SubscriptionLevels{},
        &UserSubscriptions{},
        &UserBookmarks{},
        ); err != nil {
        return err
    }

    return nil
}
