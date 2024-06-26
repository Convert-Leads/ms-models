package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	OrganisationId uint   `json:"organisation_id"`
	Name           string `json:"name"`
	BannerID       *uint
	Banner         *Media        `gorm:"foreignKey:BannerID" json:"banner"`
	Description    string        `json:"description"`
	Tags           []MetadataTag `gorm:"many2many:group_tags;" json:"tags"`
	AllowedUsers   []User        `gorm:"many2many:group_allowed_users;" json:"allowed_users"`
	Members        []User        `gorm:"many2many:user_groups;" json:"members"`
	Posts          []Post        `gorm:"foreignKey:GroupId" json:"posts"`
}
