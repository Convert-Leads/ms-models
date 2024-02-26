package models

import "gorm.io/gorm"

type SubscriptionACL struct {
    gorm.Model
    SubscriptionLevelId uint   `json:"subscription_level_id"`
    Name                string `json:"name"`
    Description         string `json:"description"`
    AllowedMetadataTags []MetadataTag `gorm:"many2many:subscription_metadata_tags;"`
    AllowedContentTypes []ContentType `gorm:"many2many:subscription_content_types;"`
}