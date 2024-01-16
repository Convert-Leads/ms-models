package models

import (
    "gorm.io/gorm"
    "time"
    "database/sql"
)

type Content struct {
    gorm.Model
    OrganisationId  uint        `json:"organisation_id"`
    ContentTypeId   uint        `json:"content_type_id"`
    AvailableDatetime *time.Time `json:"available_datetime"`
    ExpiryDatetime  *time.Time `json:"expiry_datetime"`
    Pinned          bool       `json:"pinned"`
    MediaId         uint        `json:"media_id"`
}
