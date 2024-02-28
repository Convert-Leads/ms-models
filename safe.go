package models

import (
	"gorm.io/gorm"
)

type Safe struct {
	gorm.Model
	Name         string         `gorm:"unique;not null" json:"name"`
	Image        string         `json:"image"` // Assuming this is a URL or identifier for an image
	MetadataTags []*MetadataTag `gorm:"many2many:safe_metadata_tags;" json:"tags"` // Many2Many relationship
	Content 		 []*Content     `gorm:"many2many:safe_content;" json:"content"`             // Many2Many relationship
}
