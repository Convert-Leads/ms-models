package models

import (
    "gorm.io/gorm"
)

type ContentMetadataTag struct {
    gorm.Model
    ContentId     int `json:"content_id"`
    MetadataTagId int `json:"metadata_tag_id"`
}
