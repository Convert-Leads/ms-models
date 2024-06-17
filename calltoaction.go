package models

import (
    "gorm.io/gorm"
)

type CallToAction struct {
	gorm.Model
	Text 			 string `json:"text"`
	Type       string `json:"type"`
	ContentID uint   `json:"content_id"`
	Content   *Content `json:"content"`
	CollectionID uint   `json:"collection_id"`
	Collection   *Collection `json:"collection"`
	URL        string `json:"url"`
}