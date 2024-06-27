package models

import (
    "gorm.io/gorm"
)

type SubscriptionLevel struct {
    gorm.Model
    Name           string             `json:"name"`
    MediaID        *uint               `json:"media_id"`
    Media          *Media              `gorm:"polymorphic:Parent;polymorphicValue:subscriptionlevels;foreignKey:MediaID" json:"media"`
    Description    string             `json:"description"`
    PriceOptions   []PriceOption      `json:"price_options" gorm:"foreignKey:SubscriptionLevelId"`
    OrganisationId uint               `json:"organisation_id"`
    AllowedMetadataTags []MetadataTag `gorm:"many2many:subscription_metadata_tags;" json:"allowed_metadata_tags" `
    AllowedContentTypes []ContentType `gorm:"many2many:subscription_content_types;" json:"allowed_content_types"`
    Subscribers    []UserSubscription `json:"subscribers" gorm:"foreignKey:SubscriptionLevelId"`
    SubscriberCount int                `gorm:"-" json:"subscriber_count"`
    Active         bool               `json:"active" gorm:"default:true"`
}
