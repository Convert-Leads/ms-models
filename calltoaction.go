package models

import (
    "gorm.io/gorm"
)

type CallToAction struct {
	gorm.Model
	Text 			 string `json:"text"`
	Type       string `json:"type"`
	ParentID uint   `json:"parent_id"`
	ParentType string `json:"parent_type"`
	TargetContentID *uint   `json:"target_content_id"`
	Content   *Content `gorm:"foreignKey:TargetContentID" json:"content"`
	CollectionID *uint   `json:"collection_id"`
	Collection   *Collection `gorm:"foreignKey:CollectionID" json:"collection"`
	URL        string `json:"url"`
}