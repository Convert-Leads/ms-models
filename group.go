package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	OrganisationId uint          `json:"organisation_id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Tags           []MetadataTag `gorm:"many2many:content_tags;" json:"tags"`
	AllowedUsers   []User        `gorm:"many2many:content_allowed_users;" json:"allowed_users"`
	Members        []User        `gorm:"many2many:user_groups;" json:"members"`
	Posts          []Post        `gorm:"foreignKey:GroupId" json:"posts"`
}
