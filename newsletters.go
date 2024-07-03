package models

import (
	"gorm.io/gorm"
)

type NewsletterTemplate struct {
	gorm.Model
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Elements    []NewsletterElement `json:"elements" gorm:"polymorphic:Parent;polymorphicValue:NewsletterTemplates"`
}

type Newsletter struct {
	gorm.Model
	OrganisationID uint              `json:"organisation_id"`
	TemplateID     uint              `json:"template_id"`
	ContentID      uint              `json:"content_id"`
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	Elements       []NewsletterElement `json:"elements" gorm:"polymorphic:Parent;polymorphicValue:Newsletters"`
}

type NewsletterElement struct {
	gorm.Model
	ParentID       uint                         `json:"parent_id"`
	ParentType     string                       `json:"parent_type" gorm:"type:varchar(50);polymorphic:Parent"`
	Type           string                       `json:"type"`
	Data           string                       `json:"data"`
	Order          int                          `json:"order"`
	MediaID        *uint                        `json:"media_id"`
	Media          *Media                       `gorm:"polymorphic:Parent;polymorphicValue:NewsletterElements;foreignKey:MediaID" json:"media"`
	Decorations    NewsletterElementDecoration  `json:"decorations" gorm:"foreignKey:NewsletterElementID"`
	Children       []NewsletterElement          `json:"children" gorm:"foreignKey:ParentID"` // Reference to nested elements
}

type NewsletterElementDecoration struct {
	gorm.Model
	NewsletterElementID uint   `json:"newsletter_element_id"`
	BgColor             string `json:"bg_color"`
	FontSize            int    `json:"font_size"`
	FontWeight          string `json:"font_weight"`
	FontFamily          string `json:"font_family"`
	Alignment           string `json:"alignment"`
	Padding             string `json:"padding"`
	Margin              string `json:"margin"`
}
