package models

import (
    "gorm.io/gorm"
    "time"
)

type Content struct {
    gorm.Model
    OrganisationId     uint       `json:"organisation_id"`
    ContentTypeId      uint       `json:"content_type_id"`
    AvailableDatetime  *time.Time `json:"available_datetime"`
    ExpiryDatetime     *time.Time `json:"expiry_datetime"`
    Pinned             bool       `json:"pinned"`
    Media              Media      `gorm:"foreignKey:ContentId" json:"media"` // Association to Media with JSON annotation
    UserInteractions   []UserInteraction `gorm:"foreignKey:ContentId" json:"user_interactions"` // Association to UserInteraction with JSON annotation
    DeltaContent       DeltaContent `gorm:"foreignKey:ContentId" json:"delta_content"` // Association to DeltaContent with JSON annotation
    CreatedBy  uint   `json:"created_by"`  // ID of the user who created the record
    UpdatedBy  uint   `json:"updated_by"`  // ID of the user who last updated the record
}
