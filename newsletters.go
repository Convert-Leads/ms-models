package models

import (
	"gorm.io/gorm"
)

type NewsletterTemplate struct {
	gorm.Model
	Title       string              `json:"title"`
	Description string              `json:"description"`
	ThumbnailID        *uint       
  Thumbnail          *Media     `gorm:"polymorphic:Parent;polymorphicValue:NewsletterTemplates;foreignKey:ThumbnailID" json:"thumbnail"` 
	Elements    []NewsletterElement `json:"elements" gorm:"polymorphic:Parent;polymorphicValue:NewsletterTemplates"`
}

type Newsletter struct {
	gorm.Model
	OrganisationID uint              `json:"organisation_id"`
	TemplateID     uint              `json:"template_id"`
	ContentID      uint              `json:"content_id"`
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	ThumbnailID        *uint       
  Thumbnail          *Media     `gorm:"polymorphic:Parent;polymorphicValue:NewsletterTemplates;foreignKey:ThumbnailID" json:"thumbnail"` 
	Elements       []NewsletterElement `json:"elements" gorm:"polymorphic:Parent;polymorphicValue:Newsletters"`
}

type NewsletterElement struct {
	gorm.Model
	ParentID       uint                         `json:"parent_id"`
	ParentType     string                       `json:"parent_type" gorm:"type:varchar(50);polymorphic:Parent"`
	Type           string                       `json:"type"`
	Data           string                       `json:"data"`
	Order          int                          `json:"order"`
	Media          *Media                       `gorm:"polymorphic:Parent;polymorphicValue:NewsletterElements" json:"media"`
	Decorations    NewsletterElementDecoration  `json:"decorations" gorm:"foreignKey:NewsletterElementID"`
	Children       []NewsletterElement          `json:"children" gorm:"foreignKey:ParentID;polymorphicValue:NewsletterElements"` // Reference to nested elements
	CTA						*CallToAction                `json:"cta" gorm:"polymorphic:Parent;polymorphicValue:NewsletterElements"`
}


type NewsletterElementDecoration struct {
	gorm.Model
	NewsletterElementID uint   `json:"newsletter_element_id"`
	BgColor             string `json:"bg_color"`
	FontSize            int    `json:"font_size"`
	FontWeight          string `json:"font_weight"`
	FontFamily          string `json:"font_family"`
	FontColour					string `json:"font_colour"`
	Alignment           string `json:"alignment"`
	Padding             string `json:"padding"`
	Margin              string `json:"margin"`
	Height              string `json:"height"`
	Width               string `json:"width"`
}
