package models

import (
    "gorm.io/gorm"
)

type CallToAction struct {
	gorm.Model
	Text 			 string `json:"text"`
	Type       string `json:"type"`
	ContentID uint   `json:"content_id"`
	TargetContentID uint   `json:"target_content_id"`
	Content   *Content `gorm:"foreignKey:TargetContentID" json:"content"`
	CollectionID uint   `json:"collection_id"`
	Collection   *Collection `gorm:"foreignKey:CollectionID" json:"collection"`
	URL        string `json:"url"`
}