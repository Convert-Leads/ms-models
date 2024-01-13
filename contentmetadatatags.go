package models

import (
    "gorm.io/gorm"
)

type ContentMetadataTags struct {
    gorm.Model
    ContentId     int `json:"content_id"`
    MetadataTagId int `json:"metadata_tag_id"`
}
