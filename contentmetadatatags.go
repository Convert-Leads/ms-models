package models

import (
    "gorm.io/gorm"
)

type ContentMetadataTag struct {
    gorm.Model
    ContentId     uint `json:"content_id"`
    MetadataTagId uint `json:"metadata_tag_id"`
}
