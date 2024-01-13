package models

import (
    "gorm.io/gorm"
)

type Content struct {
    gorm.Model
    OrganisationId  int        `json:"organisation_id"`
    ContentTypeId   int        `json:"content_type_id"`
    AvailableDatetime *time.Time `json:"available_datetime"`
    ExpiryDatetime  *time.Time `json:"expiry_datetime"`
    Pinned          bool       `json:"pinned"`
    MediaId         int        `json:"media_id"`
}
