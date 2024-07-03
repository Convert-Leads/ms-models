
package models

import (
	"gorm.io/gorm"
)

type NewsletterTemplate struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Elements		[]NewsletterElement `json:"elements" gorm:"foreignKey:ParentID;polymorphic:Parent;polymorphicValue:newsletter_templates"`
}

type Newsletter struct {
	gorm.Model
	OrganisationID uint `json:"organisation_id"`
	TemplateID     uint `json:"template_id"`
	ContentID      uint `json:"content_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Elements		[]NewsletterElement `json:"elements" gorm:"foreignKey:ParentID;polymorphic:Parent;polymorphicValue:newsletters"`
}

type NewsletterElement struct {
	gorm.Model
	ParentID     uint   `json:"parent_id"`
	ParentType   string `json:"parent_type"`
	Type         string `json:"type"`
	Data         string `json:"data"`
	Order        int    `json:"order"`
	MediaID      *uint  `json:"media_id"`
	Media        *Media `gorm:"polymorphic:Parent;polymorphicValue:newsletter_elements;foreignKey:MediaID" json:"media"`
	Decorations  NewsletterElementDecoration `json:"decorations" gorm:"foreignKey:NewsletterElementID"`
}

type NewsletterElementDecoration struct {
	gorm.Model
	NewsletterElementID uint   `json:"newsletter_element_id"`
	BgColor             string `json:"bg_color"`
	FontSize            int    `json:"font_size"`
	FontWeight          string `json:"font_weight"`
	FontFamily					string `json:"font_family"`
	Alignment           string `json:"alignment"`
	Padding             string `json:"padding"`
	Margin              string `json:"margin"`
}

