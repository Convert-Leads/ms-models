package models

import (
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB, models []interface{}) error {
    // AutoMigrate provided models
    if err := db.AutoMigrate(models...); err != nil {
        return err
    }

    return nil
}
